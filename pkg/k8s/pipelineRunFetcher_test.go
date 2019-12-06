package k8s

import (
	"testing"

	api "github.com/SAP/stewardci-core/pkg/apis/steward/v1alpha1"
	stewardLister "github.com/SAP/stewardci-core/pkg/client/listers/steward/v1alpha1"
	"github.com/SAP/stewardci-core/pkg/k8s/fake"
	"gotest.tools/assert"
	"k8s.io/client-go/tools/cache"
)

func Test_ClientBasedPipelineRunFetcher_ByName_NotExisting(t *testing.T) {
	// SETUP
	factory := fake.NewClientFactory()
	client := factory.StewardV1alpha1()
	examinee := NewClientBasedPipelineRunFetcher(client)

	// EXERCISE
	pipelineRun, resultErr := examinee.ByName(ns1, "NotExisting1")

	// VERIFY
	assert.Assert(t, pipelineRun == nil)
	assert.NilError(t, resultErr)
}

func Test_ClientBasedPipelineRunFetcher_ByName_GoodCase(t *testing.T) {
	// SETUP
	const (
		secretName = "secret1"
	)
	run := newPipelineRunWithSecret(ns1, run1, secretName)
	factory := fake.NewClientFactory(run)
	client := factory.StewardV1alpha1()
	examinee := NewClientBasedPipelineRunFetcher(client)

	// EXERCISE
	resultObj, resultErr := examinee.ByName(ns1, run1)

	// VERIFY
	assert.NilError(t, resultErr)
	assert.DeepEqual(t, run, resultObj)
}

func Test_ClientBasedPipelineRunFetcher_ByKey_GoodCase(t *testing.T) {
	// SETUP
	const (
		secretName = "secret1"
	)
	run := newPipelineRunWithSecret(ns1, run1, secretName)
	factory := fake.NewClientFactory(run)
	client := factory.StewardV1alpha1()
	key := fake.ObjectKey(run1, ns1)
	examinee := NewClientBasedPipelineRunFetcher(client)

	// EXERCISE
	resultObj, resultErr := examinee.ByKey(key)

	// VERIFY
	assert.NilError(t, resultErr)
	assert.DeepEqual(t, run, resultObj)
}

func Test_ListerBasedPipelineRunFetcher_ByName_NotExisting(t *testing.T) {
	// SETUP
	lister := createLister()
	examinee := NewListerBasedPipelineRunFetcher(lister)

	// EXERCISE
	pipelineRun, resultErr := examinee.ByName(ns1, "NotExisting1")

	// VERIFY
	assert.Assert(t, pipelineRun == nil)
	assert.NilError(t, resultErr)
}

func Test_ListerBasedPipelineRunFetcher_ByName_GoodCase(t *testing.T) {
	// SETUP
	const (
		secretName = "secret1"
	)
	run := newPipelineRunWithSecret(ns1, run1, secretName)
	lister := createLister(run)
	examinee := NewListerBasedPipelineRunFetcher(lister)
	// EXERCISE
	resultObj, resultErr := examinee.ByName(ns1, run1)

	// VERIFY
	assert.NilError(t, resultErr)
	assert.DeepEqual(t, run, resultObj)
}

func Test_ListerBasedPipelineRunFetcher_ByKey_GoodCase(t *testing.T) {
	// SETUP
	const (
		secretName = "secret1"
	)
	run := newPipelineRunWithSecret(ns1, run1, secretName)
	lister := createLister(run)
	examinee := NewListerBasedPipelineRunFetcher(lister)
	key := fake.ObjectKey(run1, ns1)

	// EXERCISE
	resultObj, resultErr := examinee.ByKey(key)

	// VERIFY
	assert.NilError(t, resultErr)
	assert.DeepEqual(t, run, resultObj)
}

func createLister(runs ...*api.PipelineRun) stewardLister.PipelineRunLister {
	indexer := cache.NewIndexer(cache.MetaNamespaceKeyFunc,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
	for _, run := range runs {
		indexer.Add(run)
	}
	return stewardLister.NewPipelineRunLister(indexer)
}
