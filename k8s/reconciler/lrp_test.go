package reconciler_test

import (
	"context"

	"code.cloudfoundry.org/eirini"
	"code.cloudfoundry.org/eirini/api"
	"code.cloudfoundry.org/eirini/k8s/reconciler"
	"code.cloudfoundry.org/eirini/k8s/reconciler/reconcilerfakes"
	eiriniv1 "code.cloudfoundry.org/eirini/pkg/apis/eirini/v1"
	eiriniv1scheme "code.cloudfoundry.org/eirini/pkg/generated/clientset/versioned/scheme"
	"code.cloudfoundry.org/lager/lagertest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	appsv1 "k8s.io/api/apps/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var _ = Describe("reconciler.LRP", func() {
	var (
		logger            *lagertest.TestLogger
		controllerClient  *reconcilerfakes.FakeClient
		statusClient      *reconcilerfakes.FakeStatusWriter
		statefulsetGetter *reconcilerfakes.FakeStatefulSetGetter
		desirer           *reconcilerfakes.FakeLRPDesirer
		scheme            *runtime.Scheme
		lrpreconciler     *reconciler.LRP
		resultErr         error
	)

	BeforeEach(func() {
		controllerClient = new(reconcilerfakes.FakeClient)
		desirer = new(reconcilerfakes.FakeLRPDesirer)
		statusClient = new(reconcilerfakes.FakeStatusWriter)
		statefulsetGetter = new(reconcilerfakes.FakeStatefulSetGetter)
		logger = lagertest.NewTestLogger("lrp-reconciler")
		scheme = eiriniv1scheme.Scheme
		lrpreconciler = reconciler.NewLRP(logger, controllerClient, desirer, statefulsetGetter, scheme)

		controllerClient.GetStub = func(c context.Context, nn types.NamespacedName, o client.Object) error {
			lrp, ok := o.(*eiriniv1.LRP)
			Expect(ok).To(BeTrue())
			lrp.Name = "some-lrp"
			lrp.Namespace = "some-ns"
			lrp.Spec.GUID = "the-lrp-guid"
			lrp.Spec.Version = "the-lrp-version"
			lrp.Spec.Command = []string{"ls", "-la"}
			lrp.Spec.Instances = 10
			lrp.Spec.ProcessType = "web"
			lrp.Spec.AppName = "the-app"
			lrp.Spec.AppGUID = "the-app-guid"
			lrp.Spec.OrgName = "the-org"
			lrp.Spec.OrgGUID = "the-org-guid"
			lrp.Spec.SpaceName = "the-space"
			lrp.Spec.SpaceGUID = "the-space-guid"
			lrp.Spec.Image = "eirini/dorini"
			lrp.Spec.Env = map[string]string{
				"FOO": "BAR",
			}
			lrp.Spec.Ports = []int32{8080, 9090}
			lrp.Spec.MemoryMB = 1024
			lrp.Spec.DiskMB = 512
			lrp.Spec.CPUWeight = 128
			lrp.Spec.LastUpdated = "now"
			lrp.Spec.AppRoutes = []eiriniv1.Route{
				{Hostname: "foo.io", Port: 8080}, {Hostname: "bar.io", Port: 9090},
			}
			lrp.Spec.Sidecars = []eiriniv1.Sidecar{
				{
					Name:     "hello-sidecar",
					Command:  []string{"sh", "-c", "echo hello"},
					MemoryMB: 8,
					Env: map[string]string{
						"SIDE": "BUS",
					},
				},
				{
					Name:     "bye-sidecar",
					Command:  []string{"sh", "-c", "echo bye"},
					MemoryMB: 16,
					Env: map[string]string{
						"SIDE": "CAR",
					},
				},
			}
			lrp.Spec.VolumeMounts = []eiriniv1.VolumeMount{
				{
					MountPath: "/path/to/mount",
					ClaimName: "claim-q1",
				},
				{
					MountPath: "/path/in/the/other/direction",
					ClaimName: "claim-c2",
				},
			}
			lrp.Spec.Health = eiriniv1.Healthcheck{
				Type:      "http",
				Port:      9090,
				Endpoint:  "/heath",
				TimeoutMs: 80,
			}

			lrp.Spec.UserDefinedAnnotations = map[string]string{
				"user-annotaions.io": "yes",
			}

			return nil
		}
		controllerClient.StatusReturns(statusClient)
		statefulsetGetter.GetReturns(&appsv1.StatefulSet{}, nil)
		desirer.GetReturns(nil, eirini.ErrNotFound)
	})

	JustBeforeEach(func() {
		_, resultErr = lrpreconciler.Reconcile(context.Background(), reconcile.Request{
			NamespacedName: types.NamespacedName{
				Namespace: "some-ns",
				Name:      "app",
			},
		})
	})

	It("creates a statefulset for each CRD", func() {
		Expect(resultErr).NotTo(HaveOccurred())

		Expect(desirer.UpdateCallCount()).To(Equal(0))
		Expect(desirer.DesireCallCount()).To(Equal(1))

		_, ns, lrp, _ := desirer.DesireArgsForCall(0)
		Expect(ns).To(Equal("some-ns"))
		Expect(lrp.GUID).To(Equal("the-lrp-guid"))
		Expect(lrp.Version).To(Equal("the-lrp-version"))
		Expect(lrp.Command).To(ConsistOf("ls", "-la"))
		Expect(lrp.TargetInstances).To(Equal(10))
		Expect(lrp.PrivateRegistry).To(BeNil())
		Expect(lrp.ProcessType).To(Equal("web"))
		Expect(lrp.AppName).To(Equal("the-app"))
		Expect(lrp.AppGUID).To(Equal("the-app-guid"))
		Expect(lrp.OrgName).To(Equal("the-org"))
		Expect(lrp.OrgGUID).To(Equal("the-org-guid"))
		Expect(lrp.SpaceName).To(Equal("the-space"))
		Expect(lrp.SpaceGUID).To(Equal("the-space-guid"))
		Expect(lrp.Image).To(Equal("eirini/dorini"))
		Expect(lrp.Env).To(Equal(map[string]string{
			"FOO": "BAR",
		}))
		Expect(lrp.Ports).To(Equal([]int32{8080, 9090}))
		Expect(lrp.MemoryMB).To(Equal(int64(1024)))
		Expect(lrp.DiskMB).To(Equal(int64(512)))
		Expect(lrp.CPUWeight).To(Equal(uint8(128)))
		Expect(lrp.LastUpdated).To(Equal("now"))
		Expect(lrp.AppURIs).To(Equal([]api.Route{
			{Hostname: "foo.io", Port: 8080},
			{Hostname: "bar.io", Port: 9090},
		}))
		Expect(lrp.AppURIs).To(ConsistOf(
			api.Route{Hostname: "foo.io", Port: 8080},
			api.Route{Hostname: "bar.io", Port: 9090},
		))
		Expect(lrp.Sidecars).To(Equal([]api.Sidecar{
			{
				Name:     "hello-sidecar",
				Command:  []string{"sh", "-c", "echo hello"},
				MemoryMB: 8,
				Env: map[string]string{
					"SIDE": "BUS",
				},
			},
			{
				Name:     "bye-sidecar",
				Command:  []string{"sh", "-c", "echo bye"},
				MemoryMB: 16,
				Env: map[string]string{
					"SIDE": "CAR",
				},
			},
		}))
		Expect(lrp.VolumeMounts).To(Equal([]api.VolumeMount{
			{
				MountPath: "/path/to/mount",
				ClaimName: "claim-q1",
			},
			{
				MountPath: "/path/in/the/other/direction",
				ClaimName: "claim-c2",
			},
		}))
		Expect(lrp.Health).To(Equal(api.Healthcheck{
			Type:      "http",
			Port:      9090,
			Endpoint:  "/heath",
			TimeoutMs: 80,
		}))
		Expect(lrp.UserDefinedAnnotations).To(Equal(map[string]string{
			"user-annotaions.io": "yes",
		}))
	})

	It("sets an owner reference in the statefulset", func() {
		Expect(resultErr).NotTo(HaveOccurred())

		Expect(desirer.DesireCallCount()).To(Equal(1))
		_, _, _, setOwnerFns := desirer.DesireArgsForCall(0)
		Expect(setOwnerFns).To(HaveLen(1))
		setOwnerFn := setOwnerFns[0]

		st := &appsv1.StatefulSet{ObjectMeta: v1.ObjectMeta{Namespace: "some-ns"}}
		Expect(setOwnerFn(st)).To(Succeed())
		Expect(st.ObjectMeta.OwnerReferences).To(HaveLen(1))
		Expect(st.ObjectMeta.OwnerReferences[0].Kind).To(Equal("LRP"))
		Expect(st.ObjectMeta.OwnerReferences[0].Name).To(Equal("some-lrp"))
	})

	When("private registry credentials are specified in the LRP CRD", func() {
		BeforeEach(func() {
			controllerClient.GetStub = func(c context.Context, nn types.NamespacedName, o client.Object) error {
				lrp, ok := o.(*eiriniv1.LRP)
				Expect(ok).To(BeTrue())

				lrp.Spec.Image = "private-registry.com:5000/repo/app-image:latest"
				lrp.Spec.PrivateRegistry = &eiriniv1.PrivateRegistry{
					Username: "docker-user",
					Password: "docker-password",
				}

				return nil
			}
		})

		It("configures a private registry", func() {
			_, _, lrp, _ := desirer.DesireArgsForCall(0)
			privateRegistry := lrp.PrivateRegistry
			Expect(privateRegistry).NotTo(BeNil())
			Expect(privateRegistry.Username).To(Equal("docker-user"))
			Expect(privateRegistry.Password).To(Equal("docker-password"))
			Expect(privateRegistry.Server).To(Equal("private-registry.com"))
		})

		When("the image URL does not contain an image registry host", func() {
			BeforeEach(func() {
				controllerClient.GetStub = func(c context.Context, nn types.NamespacedName, o client.Object) error {
					lrp, ok := o.(*eiriniv1.LRP)
					Expect(ok).To(BeTrue())

					lrp.Spec.Image = "eirini/dorini"
					lrp.Spec.PrivateRegistry = &eiriniv1.PrivateRegistry{
						Username: "docker-user",
						Password: "docker-password",
					}

					return nil
				}
			})

			It("configures the private registry server with the dockerhub host", func() {
				_, _, lrp, _ := desirer.DesireArgsForCall(0)
				Expect(lrp.PrivateRegistry.Server).To(Equal("index.docker.io/v1/"))
			})
		})
	})

	When("a CRD for this app already exists", func() {
		BeforeEach(func() {
			desirer.GetReturns(nil, nil)
		})

		It("updates it", func() {
			Expect(resultErr).NotTo(HaveOccurred())

			Expect(desirer.UpdateCallCount()).To(Equal(1))
			_, lrp := desirer.UpdateArgsForCall(0)
			Expect(lrp.TargetInstances).To(Equal(10))
			Expect(lrp.AppURIs).To(ConsistOf(
				api.Route{Hostname: "foo.io", Port: 8080},
				api.Route{Hostname: "bar.io", Port: 9090},
			))
		})
	})

	When("an app instance becomes unready", func() {
		var statusWriter *reconcilerfakes.FakeStatusWriter

		BeforeEach(func() {
			statusWriter = new(reconcilerfakes.FakeStatusWriter)
			controllerClient.StatusReturns(statusWriter)

			desirer.GetReturns(nil, nil)
			statefulsetGetter.GetReturns(&appsv1.StatefulSet{Status: appsv1.StatefulSetStatus{ReadyReplicas: 9}}, nil)
		})

		It("the CRD status is updated accordingly", func() {
			Expect(resultErr).NotTo(HaveOccurred())

			Expect(statusWriter.UpdateCallCount()).To(Equal(1))
			_, obj, _ := statusWriter.UpdateArgsForCall(0)
			lrp, ok := obj.(*eiriniv1.LRP)
			Expect(ok).To(BeTrue())
			Expect(lrp.Status).To(Equal(eiriniv1.LRPStatus{
				Replicas: 9,
			}))
		})

		When("statefulset getter fails to get the statefulset", func() {
			BeforeEach(func() {
				statefulsetGetter.GetReturns(nil, errors.New("boom"))
			})

			It("does not update the statefulset status", func() {
				Expect(resultErr).To(MatchError(ContainSubstring("boom")))

				Expect(statusWriter.UpdateCallCount()).To(Equal(0))
				Expect(desirer.UpdateCallCount()).To(Equal(1))
			})
		})

		When("the controller client fails to update the LRP status", func() {
			BeforeEach(func() {
				statusWriter.UpdateReturns(errors.New("bom"))
			})

			It("does not update the statefulset status", func() {
				Expect(resultErr).To(MatchError(ContainSubstring("bom")))

				Expect(statusWriter.UpdateCallCount()).To(Equal(1))
				Expect(desirer.UpdateCallCount()).To(Equal(1))
			})
		})
	})

	When("the LRP doesn't exist", func() {
		BeforeEach(func() {
			controllerClient.GetStub = func(c context.Context, nn types.NamespacedName, o client.Object) error {
				return apierrors.NewNotFound(schema.GroupResource{}, "my-lrp")
			}
		})

		It("does not return an error", func() {
			Expect(resultErr).NotTo(HaveOccurred())
		})
	})

	When("the controller client fails to get the CRD", func() {
		BeforeEach(func() {
			controllerClient.GetStub = func(c context.Context, nn types.NamespacedName, o client.Object) error {
				return errors.New("boom")
			}
		})

		It("returns an error", func() {
			Expect(resultErr).To(MatchError(ContainSubstring("boom")))
		})
	})

	When("the lrp desirer fails to get the app", func() {
		BeforeEach(func() {
			desirer.GetReturns(nil, errors.New("boom"))
		})

		It("returns an error", func() {
			Expect(resultErr).To(MatchError("failed to get lrp: boom"))
		})
	})

	When("the lrp desirer fails to desire the app", func() {
		BeforeEach(func() {
			desirer.DesireReturns(errors.New("boom"))
		})

		It("returns an error", func() {
			Expect(resultErr).To(MatchError("failed to desire lrp: boom"))
		})
	})

	When("the lrp desirer fails to update the app", func() {
		BeforeEach(func() {
			desirer.GetReturns(nil, nil)
			desirer.UpdateReturns(errors.New("boom"))
		})

		It("returns an error", func() {
			Expect(resultErr).To(MatchError(ContainSubstring("boom")))
		})
	})
})
