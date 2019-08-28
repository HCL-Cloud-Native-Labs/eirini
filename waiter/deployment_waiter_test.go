package waiter_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"

	. "code.cloudfoundry.org/eirini/waiter"
	"code.cloudfoundry.org/eirini/waiter/waiterfakes"
	"code.cloudfoundry.org/lager/lagertest"
)

var _ = Describe("DeploymentWaiter", func() {

	var (
		fakeLister *waiterfakes.FakeDeploymentLister
		waiter     Deployment
		logger     *lagertest.TestLogger
	)

	BeforeEach(func() {
		fakeLister = new(waiterfakes.FakeDeploymentLister)
		logger = lagertest.NewTestLogger("test")
		waiter = Deployment{
			Deployments:       fakeLister,
			Logger:            logger,
			ListLabelSelector: "my=label",
		}
	})

	It("reports ready if all Deployments are updated and ready", func() {
		ssList := &appsv1.DeploymentList{
			Items: []appsv1.Deployment{createDeployment(replicaStatuses{
				desired:     1,
				ready:       1,
				updated:     1,
				available:   1,
				unavailable: 0,
			})},
		}
		fakeLister.ListReturns(ssList, nil)

		Expect(waiter.IsReady()).To(BeTrue())
	})

	It("reports not ready if observed generation is not updated yet", func() {
		outdatedSS := createDeployment(replicaStatuses{
			desired:     1,
			ready:       1,
			updated:     1,
			available:   1,
			unavailable: 0,
		})
		outdatedSS.Generation = 3
		outdatedSS.Status.ObservedGeneration = 2

		fakeLister.ListReturns(&appsv1.DeploymentList{Items: []appsv1.Deployment{outdatedSS}}, nil)

		Expect(waiter.IsReady()).To(BeFalse())
	})

	It("reports not ready if there are non-ready replicas", func() {
		ssList := &appsv1.DeploymentList{
			Items: []appsv1.Deployment{createDeployment(replicaStatuses{
				desired: 3,
				ready:   1,
			})},
		}
		fakeLister.ListReturns(ssList, nil)

		Expect(waiter.IsReady()).To(BeFalse())
	})

	It("reports not ready if there are non-updated replicas", func() {
		ssList := &appsv1.DeploymentList{
			Items: []appsv1.Deployment{createDeployment(replicaStatuses{
				desired: 3,
				ready:   3,
				updated: 2,
			})},
		}
		fakeLister.ListReturns(ssList, nil)

		Expect(waiter.IsReady()).To(BeFalse())
	})

	It("reports not ready if current replicas are less than desired", func() {
		ssList := &appsv1.DeploymentList{
			Items: []appsv1.Deployment{createDeployment(replicaStatuses{
				desired:   3,
				ready:     3,
				updated:   3,
				available: 1,
			})},
		}
		fakeLister.ListReturns(ssList, nil)

		Expect(waiter.IsReady()).To(BeFalse())
	})

	It("reports not ready is there remain unavailable replicas", func() {
		ssList := &appsv1.DeploymentList{
			Items: []appsv1.Deployment{createDeployment(replicaStatuses{
				desired:     3,
				ready:       3,
				updated:     3,
				available:   3,
				unavailable: 1,
			})},
		}
		fakeLister.ListReturns(ssList, nil)
		Expect(waiter.IsReady()).To(BeFalse())
	})

	When("listing Deployments fails", func() {
		It("should log the error", func() {
			fakeLister.ListReturns(nil, errors.New("boom?"))
			waiter.IsReady()

			Eventually(logger.Logs, "2s").Should(HaveLen(1))
			log := logger.Logs()[0]
			Expect(log.Message).To(Equal("test.failed to list deployments"))
			Expect(log.Data).To(HaveKeyWithValue("error", "boom?"))
		})

		It("should return false", func() {
			fakeLister.ListReturns(nil, errors.New("boom?"))
			Expect(waiter.IsReady()).To(BeFalse())
		})
	})

	When("listing Deployments", func() {
		It("should use the right label selector", func() {
			ssList := &appsv1.DeploymentList{
				Items: []appsv1.Deployment{createDeployment(replicaStatuses{
					desired:   1,
					ready:     1,
					available: 1,
					updated:   1,
				})},
			}
			fakeLister.ListReturns(ssList, nil)
			waiter.IsReady()

			Expect(fakeLister.ListCallCount()).To(Equal(1))
			listOptions := fakeLister.ListArgsForCall(0)
			Expect(listOptions.LabelSelector).To(Equal("my=label"))
		})

	})
})

type replicaStatuses struct {
	desired, ready, available, updated, unavailable int32
}

func createDeployment(replicaStatuses replicaStatuses) appsv1.Deployment {
	return appsv1.Deployment{
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicaStatuses.desired,
		},
		Status: appsv1.DeploymentStatus{
			ReadyReplicas:       replicaStatuses.ready,
			UpdatedReplicas:     replicaStatuses.updated,
			AvailableReplicas:   replicaStatuses.available,
			UnavailableReplicas: replicaStatuses.unavailable,
		},
	}
}
