package task_reporter_test

import (
	"encoding/json"
	"testing"
	"time"

	"code.cloudfoundry.org/eirini/tests"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func TestStagingReporter(t *testing.T) {
	SetDefaultEventuallyTimeout(2 * time.Minute)
	RegisterFailHandler(Fail)
	RunSpecs(t, "TaskReporter Suite")
}

var (
	fixture         *tests.Fixture
	eiriniBins      tests.EiriniBinaries
	envVarOverrides []string
)

var _ = SynchronizedBeforeSuite(func() []byte {
	eiriniBins = tests.NewEiriniBinaries()
	eiriniBins.TaskReporter.Build()

	data, err := json.Marshal(eiriniBins)
	Expect(err).NotTo(HaveOccurred())

	return data
}, func(data []byte) {
	err := json.Unmarshal(data, &eiriniBins)
	Expect(err).NotTo(HaveOccurred())

	fixture = tests.NewFixture(GinkgoWriter)
})

var _ = SynchronizedAfterSuite(func() {
	fixture.Destroy()
}, func() {
	eiriniBins.TearDown()
})

var _ = BeforeEach(func() {
	envVarOverrides = []string{}
	fixture.SetUp()

	Expect(tests.CreateSecretWithStringData(fixture.Namespace, "cc-uploader-secret", fixture.Clientset, map[string]string{"foo1": "val1", "bar1": "val2"})).To(Succeed())
	Expect(tests.CreateSecretWithStringData(fixture.Namespace, "eirini-client-secret", fixture.Clientset, map[string]string{"foo2": "val1", "bar2": "val2"})).To(Succeed())
	Expect(tests.CreateSecretWithStringData(fixture.Namespace, "ca-cert-secret", fixture.Clientset, map[string]string{"foo3": "val1", "bar3": "val2"})).To(Succeed())
})

var _ = AfterEach(func() {
	fixture.TearDown()
})
