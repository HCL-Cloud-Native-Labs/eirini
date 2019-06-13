package event_test

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"code.cloudfoundry.org/eirini/events"
	. "code.cloudfoundry.org/eirini/k8s/informers/event"
	"code.cloudfoundry.org/eirini/models/cf"
	"code.cloudfoundry.org/lager/lagertest"
	"code.cloudfoundry.org/runtimeschema/cc_messages"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/testing"
	testcore "k8s.io/client-go/testing"
)

var _ = Describe("Event", func() {

	const namespace = "milkyway"

	var (
		client *fake.Clientset

		reportChan      chan events.CrashReport
		informerStopper chan struct{}

		watcher *watch.FakeWatcher

		logger *lagertest.TestLogger

		crashTime  meta.Time
		informerWG sync.WaitGroup
	)

	BeforeEach(func() {
		reportChan = make(chan events.CrashReport)
		informerStopper = make(chan struct{})

		logger = lagertest.NewTestLogger("crash-event-logger-test")
		client = fake.NewSimpleClientset()
		crashInformer := NewCrashInformer(client, 0, namespace, reportChan, informerStopper, logger)

		watcher = watch.NewFake()
		client.PrependWatchReactor("pods", testing.DefaultWatchReactor(watcher, nil))

		informerWG.Add(1)
		go func() {
			crashInformer.Start()
			informerWG.Done()
		}()
	})

	AfterEach(func() {
		close(informerStopper)
		informerWG.Wait()
	})

	Context("When app is in CrashLoopBackOff", func() {
		BeforeEach(func() {
			normy := createPod()
			watcher.Add(normy)

			loopy := createPod()
			crashTime = meta.Time{Time: time.Now()}
			loopy.Status.ContainerStatuses = []v1.ContainerStatus{
				{
					RestartCount: 3,
					State: v1.ContainerState{
						Waiting: &v1.ContainerStateWaiting{
							Reason: CrashLoopBackOff,
						},
					},
					LastTerminationState: v1.ContainerState{
						Terminated: &v1.ContainerStateTerminated{
							ExitCode:  1,
							Reason:    "this describes how much you screwed up",
							StartedAt: crashTime,
						},
					},
				},
			}
			watcher.Modify(loopy)
		})

		It("should receive a crashed report", func() {
			Eventually(reportChan).Should(Receive(Equal(events.CrashReport{
				ProcessGUID: "test-pod-anno",
				AppCrashedRequest: cc_messages.AppCrashedRequest{
					Reason:          CrashLoopBackOff,
					Instance:        "test-pod-0",
					Index:           0,
					ExitStatus:      1,
					ExitDescription: "this describes how much you screwed up",
					CrashCount:      3,
					CrashTimestamp:  int64(crashTime.Time.Second()),
				},
			})))
		})

	})

	Context("When app has been terminated", func() {

		var (
			normy *v1.Pod
			termy *v1.Pod
		)

		BeforeEach(func() {
			normy = createPod()
			watcher.Add(normy)

			termy = createPod()
			crashTime = meta.Time{Time: time.Now()}
			termy.Status.ContainerStatuses = []v1.ContainerStatus{
				{
					RestartCount: 8,
					State: v1.ContainerState{
						Terminated: &v1.ContainerStateTerminated{
							Reason:    "this describes how much you screwed up",
							StartedAt: crashTime,
						},
					},
				},
			}
		})

		Context("with non-zero exit status", func() {

			BeforeEach(func() {
				termy.Status.ContainerStatuses[0].State.Terminated.ExitCode = 1
				watcher.Modify(termy)
			})

			It("should receive a crashed report", func() {
				Eventually(reportChan).Should(Receive(Equal(events.CrashReport{
					ProcessGUID: "test-pod-anno",
					AppCrashedRequest: cc_messages.AppCrashedRequest{
						Reason:          "this describes how much you screwed up",
						Instance:        "test-pod-0",
						Index:           0,
						ExitStatus:      1,
						ExitDescription: "this describes how much you screwed up",
						CrashCount:      8,
						CrashTimestamp:  int64(crashTime.Time.Second()),
					},
				})))
			})

		})

		Context("with zero exit status", func() {

			BeforeEach(func() {
				termy.Status.ContainerStatuses[0].State.Terminated.ExitCode = 0
				watcher.Modify(termy)
			})

			It("should not send reports", func() {
				Consistently(reportChan).ShouldNot(Receive())
			})

		})
	})

	Context("When a pod name is incorrect", func() {

		BeforeEach(func() {
			statelessy := createStatelessPod("test-pod")
			watcher.Add(statelessy)
			watcher.Modify(statelessy)
		})

		It("should not send reports", func() {
			Consistently(reportChan).ShouldNot(Receive())
		})

		It("should provide a helpful log message", func() {
			Consistently(reportChan).ShouldNot(Receive())

			logs := logger.Logs()
			Expect(logs).To(HaveLen(1))
			log := logs[0]
			Expect(log.Message).To(Equal("crash-event-logger-test.failed-to-parse-app-index"))
			Expect(log.Data).To(HaveKeyWithValue("pod-name", "test-pod"))
			Expect(log.Data).To(HaveKeyWithValue("guid", "test-pod-anno"))
		})

	})

	Context("When app is waiting, but NOT because of CrashLoopBackOff", func() {

		BeforeEach(func() {
			normy := createPod()
			watcher.Add(normy)

			sleepy := createPod()
			crashTime = meta.Time{Time: time.Now()}
			sleepy.Status.ContainerStatuses = []v1.ContainerStatus{
				{
					State: v1.ContainerState{
						Waiting: &v1.ContainerStateWaiting{
							Reason: "sleepy",
						},
					},
				},
			}
			watcher.Modify(sleepy)
		})

		It("should not send reports", func() {
			Consistently(reportChan).ShouldNot(Receive())
		})

	})

	Context("When a pod has no container statuses", func() {

		Context("container statuses is nil", func() {

			BeforeEach(func() {
				normy := createPod()
				watcher.Add(normy)

				normy.Status.ContainerStatuses = nil
				watcher.Modify(normy)
			})

			It("should not send any reports", func() {
				Consistently(reportChan).ShouldNot(Receive())
			})
		})

		Context("container statuses is empty", func() {

			BeforeEach(func() {
				normy := createPod()
				watcher.Add(normy)

				normy.Status.ContainerStatuses = []v1.ContainerStatus{}
				watcher.Modify(normy)
			})

			It("should not send any reports", func() {
				Consistently(reportChan).ShouldNot(Receive())
			})
		})
	})

	Context("When pod is stopped", func() {

		BeforeEach(func() {
			event := v1.Event{
				InvolvedObject: v1.ObjectReference{
					Namespace: namespace,
					Name:      "pinky-pod",
				},
				Reason: "Killing",
			}
			_, clientErr := client.CoreV1().Events(namespace).Create(&event)
			Expect(clientErr).ToNot(HaveOccurred())

			termy := createPod()
			watcher.Add(termy)
			termy.Status.ContainerStatuses = []v1.ContainerStatus{
				{
					State: v1.ContainerState{
						Terminated: &v1.ContainerStateTerminated{
							ExitCode: 1,
						},
					},
				},
			}
			watcher.Modify(termy)
		})

		It("should not emit a crashed event", func() {
			Consistently(reportChan).ShouldNot(Receive())
		})
	})

	Context("When getting events fails", func() {
		BeforeEach(func() {
			reaction := func(action testcore.Action) (handled bool, ret runtime.Object, err error) {
				return true, nil, errors.New("boom")
			}
			client.PrependReactor("list", "events", reaction)

			termy := createPod()
			watcher.Add(termy)
			termy.Status.ContainerStatuses = []v1.ContainerStatus{
				{
					State: v1.ContainerState{
						Terminated: &v1.ContainerStateTerminated{
							ExitCode: 1,
						},
					},
				},
			}
			watcher.Modify(termy)
		})

		It("should not emit a crashed event", func() {
			Consistently(reportChan).ShouldNot(Receive())
		})

		It("should provide a helpful log message", func() {
			Consistently(reportChan).ShouldNot(Receive())

			logs := logger.Logs()
			Expect(logs).To(HaveLen(1))
			log := logs[0]
			Expect(log.Message).To(Equal("crash-event-logger-test.failed-to-get-k8s-events"))
			Expect(log.Data).To(HaveKeyWithValue("guid", "test-pod-anno"))
		})

	})
})

func createPod() *v1.Pod {
	name := "test-pod"
	return &v1.Pod{
		ObjectMeta: meta.ObjectMeta{
			Name: fmt.Sprintf("%s-%d", name, 0),
			Annotations: map[string]string{
				cf.ProcessGUID: fmt.Sprintf("%s-anno", name),
			},
			OwnerReferences: []meta.OwnerReference{
				{
					Kind: "StatefulSet",
					Name: "mr-stateful",
				},
			},
		},
		Status: v1.PodStatus{
			ContainerStatuses: []v1.ContainerStatus{
				{
					State: v1.ContainerState{
						Running: &v1.ContainerStateRunning{},
					},
				},
			},
		},
	}
}

func createStatelessPod(name string) *v1.Pod {
	return &v1.Pod{
		ObjectMeta: meta.ObjectMeta{
			Name: name,
			Annotations: map[string]string{
				cf.ProcessGUID: fmt.Sprintf("%s-anno", name),
			},
		},
		Status: v1.PodStatus{
			ContainerStatuses: []v1.ContainerStatus{
				{
					State: v1.ContainerState{
						Running: &v1.ContainerStateRunning{},
					},
				},
			},
		},
	}
}
