// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	os "os"
	sync "sync"

	commands "github.com/pivotal-cf/om/commands"
)

type ProductDownloader struct {
	DownloadProductToFileStub        func(commands.FileArtifacter, *os.File) error
	downloadProductToFileMutex       sync.RWMutex
	downloadProductToFileArgsForCall []struct {
		arg1 commands.FileArtifacter
		arg2 *os.File
	}
	downloadProductToFileReturns struct {
		result1 error
	}
	downloadProductToFileReturnsOnCall map[int]struct {
		result1 error
	}
	GetAllProductVersionsStub        func(string) ([]string, error)
	getAllProductVersionsMutex       sync.RWMutex
	getAllProductVersionsArgsForCall []struct {
		arg1 string
	}
	getAllProductVersionsReturns struct {
		result1 []string
		result2 error
	}
	getAllProductVersionsReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	GetLatestProductFileStub        func(string, string, string) (commands.FileArtifacter, error)
	getLatestProductFileMutex       sync.RWMutex
	getLatestProductFileArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	getLatestProductFileReturns struct {
		result1 commands.FileArtifacter
		result2 error
	}
	getLatestProductFileReturnsOnCall map[int]struct {
		result1 commands.FileArtifacter
		result2 error
	}
	GetLatestStemcellForProductStub        func(commands.FileArtifacter, string) (commands.StemcellArtifacter, error)
	getLatestStemcellForProductMutex       sync.RWMutex
	getLatestStemcellForProductArgsForCall []struct {
		arg1 commands.FileArtifacter
		arg2 string
	}
	getLatestStemcellForProductReturns struct {
		result1 commands.StemcellArtifacter
		result2 error
	}
	getLatestStemcellForProductReturnsOnCall map[int]struct {
		result1 commands.StemcellArtifacter
		result2 error
	}
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct {
	}
	nameReturns struct {
		result1 string
	}
	nameReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ProductDownloader) DownloadProductToFile(arg1 commands.FileArtifacter, arg2 *os.File) error {
	fake.downloadProductToFileMutex.Lock()
	ret, specificReturn := fake.downloadProductToFileReturnsOnCall[len(fake.downloadProductToFileArgsForCall)]
	fake.downloadProductToFileArgsForCall = append(fake.downloadProductToFileArgsForCall, struct {
		arg1 commands.FileArtifacter
		arg2 *os.File
	}{arg1, arg2})
	fake.recordInvocation("DownloadProductToFile", []interface{}{arg1, arg2})
	fake.downloadProductToFileMutex.Unlock()
	if fake.DownloadProductToFileStub != nil {
		return fake.DownloadProductToFileStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.downloadProductToFileReturns
	return fakeReturns.result1
}

func (fake *ProductDownloader) DownloadProductToFileCallCount() int {
	fake.downloadProductToFileMutex.RLock()
	defer fake.downloadProductToFileMutex.RUnlock()
	return len(fake.downloadProductToFileArgsForCall)
}

func (fake *ProductDownloader) DownloadProductToFileCalls(stub func(commands.FileArtifacter, *os.File) error) {
	fake.downloadProductToFileMutex.Lock()
	defer fake.downloadProductToFileMutex.Unlock()
	fake.DownloadProductToFileStub = stub
}

func (fake *ProductDownloader) DownloadProductToFileArgsForCall(i int) (commands.FileArtifacter, *os.File) {
	fake.downloadProductToFileMutex.RLock()
	defer fake.downloadProductToFileMutex.RUnlock()
	argsForCall := fake.downloadProductToFileArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ProductDownloader) DownloadProductToFileReturns(result1 error) {
	fake.downloadProductToFileMutex.Lock()
	defer fake.downloadProductToFileMutex.Unlock()
	fake.DownloadProductToFileStub = nil
	fake.downloadProductToFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *ProductDownloader) DownloadProductToFileReturnsOnCall(i int, result1 error) {
	fake.downloadProductToFileMutex.Lock()
	defer fake.downloadProductToFileMutex.Unlock()
	fake.DownloadProductToFileStub = nil
	if fake.downloadProductToFileReturnsOnCall == nil {
		fake.downloadProductToFileReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.downloadProductToFileReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ProductDownloader) GetAllProductVersions(arg1 string) ([]string, error) {
	fake.getAllProductVersionsMutex.Lock()
	ret, specificReturn := fake.getAllProductVersionsReturnsOnCall[len(fake.getAllProductVersionsArgsForCall)]
	fake.getAllProductVersionsArgsForCall = append(fake.getAllProductVersionsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetAllProductVersions", []interface{}{arg1})
	fake.getAllProductVersionsMutex.Unlock()
	if fake.GetAllProductVersionsStub != nil {
		return fake.GetAllProductVersionsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getAllProductVersionsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ProductDownloader) GetAllProductVersionsCallCount() int {
	fake.getAllProductVersionsMutex.RLock()
	defer fake.getAllProductVersionsMutex.RUnlock()
	return len(fake.getAllProductVersionsArgsForCall)
}

func (fake *ProductDownloader) GetAllProductVersionsCalls(stub func(string) ([]string, error)) {
	fake.getAllProductVersionsMutex.Lock()
	defer fake.getAllProductVersionsMutex.Unlock()
	fake.GetAllProductVersionsStub = stub
}

func (fake *ProductDownloader) GetAllProductVersionsArgsForCall(i int) string {
	fake.getAllProductVersionsMutex.RLock()
	defer fake.getAllProductVersionsMutex.RUnlock()
	argsForCall := fake.getAllProductVersionsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ProductDownloader) GetAllProductVersionsReturns(result1 []string, result2 error) {
	fake.getAllProductVersionsMutex.Lock()
	defer fake.getAllProductVersionsMutex.Unlock()
	fake.GetAllProductVersionsStub = nil
	fake.getAllProductVersionsReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *ProductDownloader) GetAllProductVersionsReturnsOnCall(i int, result1 []string, result2 error) {
	fake.getAllProductVersionsMutex.Lock()
	defer fake.getAllProductVersionsMutex.Unlock()
	fake.GetAllProductVersionsStub = nil
	if fake.getAllProductVersionsReturnsOnCall == nil {
		fake.getAllProductVersionsReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.getAllProductVersionsReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *ProductDownloader) GetLatestProductFile(arg1 string, arg2 string, arg3 string) (commands.FileArtifacter, error) {
	fake.getLatestProductFileMutex.Lock()
	ret, specificReturn := fake.getLatestProductFileReturnsOnCall[len(fake.getLatestProductFileArgsForCall)]
	fake.getLatestProductFileArgsForCall = append(fake.getLatestProductFileArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("GetLatestProductFile", []interface{}{arg1, arg2, arg3})
	fake.getLatestProductFileMutex.Unlock()
	if fake.GetLatestProductFileStub != nil {
		return fake.GetLatestProductFileStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getLatestProductFileReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ProductDownloader) GetLatestProductFileCallCount() int {
	fake.getLatestProductFileMutex.RLock()
	defer fake.getLatestProductFileMutex.RUnlock()
	return len(fake.getLatestProductFileArgsForCall)
}

func (fake *ProductDownloader) GetLatestProductFileCalls(stub func(string, string, string) (commands.FileArtifacter, error)) {
	fake.getLatestProductFileMutex.Lock()
	defer fake.getLatestProductFileMutex.Unlock()
	fake.GetLatestProductFileStub = stub
}

func (fake *ProductDownloader) GetLatestProductFileArgsForCall(i int) (string, string, string) {
	fake.getLatestProductFileMutex.RLock()
	defer fake.getLatestProductFileMutex.RUnlock()
	argsForCall := fake.getLatestProductFileArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *ProductDownloader) GetLatestProductFileReturns(result1 commands.FileArtifacter, result2 error) {
	fake.getLatestProductFileMutex.Lock()
	defer fake.getLatestProductFileMutex.Unlock()
	fake.GetLatestProductFileStub = nil
	fake.getLatestProductFileReturns = struct {
		result1 commands.FileArtifacter
		result2 error
	}{result1, result2}
}

func (fake *ProductDownloader) GetLatestProductFileReturnsOnCall(i int, result1 commands.FileArtifacter, result2 error) {
	fake.getLatestProductFileMutex.Lock()
	defer fake.getLatestProductFileMutex.Unlock()
	fake.GetLatestProductFileStub = nil
	if fake.getLatestProductFileReturnsOnCall == nil {
		fake.getLatestProductFileReturnsOnCall = make(map[int]struct {
			result1 commands.FileArtifacter
			result2 error
		})
	}
	fake.getLatestProductFileReturnsOnCall[i] = struct {
		result1 commands.FileArtifacter
		result2 error
	}{result1, result2}
}

func (fake *ProductDownloader) GetLatestStemcellForProduct(arg1 commands.FileArtifacter, arg2 string) (commands.StemcellArtifacter, error) {
	fake.getLatestStemcellForProductMutex.Lock()
	ret, specificReturn := fake.getLatestStemcellForProductReturnsOnCall[len(fake.getLatestStemcellForProductArgsForCall)]
	fake.getLatestStemcellForProductArgsForCall = append(fake.getLatestStemcellForProductArgsForCall, struct {
		arg1 commands.FileArtifacter
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetLatestStemcellForProduct", []interface{}{arg1, arg2})
	fake.getLatestStemcellForProductMutex.Unlock()
	if fake.GetLatestStemcellForProductStub != nil {
		return fake.GetLatestStemcellForProductStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getLatestStemcellForProductReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ProductDownloader) GetLatestStemcellForProductCallCount() int {
	fake.getLatestStemcellForProductMutex.RLock()
	defer fake.getLatestStemcellForProductMutex.RUnlock()
	return len(fake.getLatestStemcellForProductArgsForCall)
}

func (fake *ProductDownloader) GetLatestStemcellForProductCalls(stub func(commands.FileArtifacter, string) (commands.StemcellArtifacter, error)) {
	fake.getLatestStemcellForProductMutex.Lock()
	defer fake.getLatestStemcellForProductMutex.Unlock()
	fake.GetLatestStemcellForProductStub = stub
}

func (fake *ProductDownloader) GetLatestStemcellForProductArgsForCall(i int) (commands.FileArtifacter, string) {
	fake.getLatestStemcellForProductMutex.RLock()
	defer fake.getLatestStemcellForProductMutex.RUnlock()
	argsForCall := fake.getLatestStemcellForProductArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ProductDownloader) GetLatestStemcellForProductReturns(result1 commands.StemcellArtifacter, result2 error) {
	fake.getLatestStemcellForProductMutex.Lock()
	defer fake.getLatestStemcellForProductMutex.Unlock()
	fake.GetLatestStemcellForProductStub = nil
	fake.getLatestStemcellForProductReturns = struct {
		result1 commands.StemcellArtifacter
		result2 error
	}{result1, result2}
}

func (fake *ProductDownloader) GetLatestStemcellForProductReturnsOnCall(i int, result1 commands.StemcellArtifacter, result2 error) {
	fake.getLatestStemcellForProductMutex.Lock()
	defer fake.getLatestStemcellForProductMutex.Unlock()
	fake.GetLatestStemcellForProductStub = nil
	if fake.getLatestStemcellForProductReturnsOnCall == nil {
		fake.getLatestStemcellForProductReturnsOnCall = make(map[int]struct {
			result1 commands.StemcellArtifacter
			result2 error
		})
	}
	fake.getLatestStemcellForProductReturnsOnCall[i] = struct {
		result1 commands.StemcellArtifacter
		result2 error
	}{result1, result2}
}

func (fake *ProductDownloader) Name() string {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct {
	}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.nameReturns
	return fakeReturns.result1
}

func (fake *ProductDownloader) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *ProductDownloader) NameCalls(stub func() string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = stub
}

func (fake *ProductDownloader) NameReturns(result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *ProductDownloader) NameReturnsOnCall(i int, result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	if fake.nameReturnsOnCall == nil {
		fake.nameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.nameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *ProductDownloader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.downloadProductToFileMutex.RLock()
	defer fake.downloadProductToFileMutex.RUnlock()
	fake.getAllProductVersionsMutex.RLock()
	defer fake.getAllProductVersionsMutex.RUnlock()
	fake.getLatestProductFileMutex.RLock()
	defer fake.getLatestProductFileMutex.RUnlock()
	fake.getLatestStemcellForProductMutex.RLock()
	defer fake.getLatestStemcellForProductMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ProductDownloader) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ commands.ProductDownloader = new(ProductDownloader)
