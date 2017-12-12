package video

import (
	"github.com/drossha/vk-api-go"
)

const (
	MethodGet                = "video.get"
	MethodEdit               = "video.edit"
	MethodAdd                = "video.add"
	MethodSave               = "video.save"
	MethodDelete             = "video.delete"
	MethodRestore            = "video.restore"
	MethodSearch             = "video.search"
	MethodGetUserVideos      = "video.getUserVideos"
	MethodGetAlbums          = "video.getAlbums"
	MethodGetAlbumById       = "video.getAlbumById"
	MethodAddAlbum           = "video.addAlbum"
	MethodEditAlbum          = "video.editAlbum"
	MethodDeleteAlbum        = "video.deleteAlbum"
	MethodReorderAlbums      = "video.reorderAlbums"
	MethodReorderVideos      = "video.reorderVideos"
	MethodAddToAlbum         = "video.addToAlbum"
	MethodRemoveFromAlbum    = "video.removeFromAlbum"
	MethodGetAlbumsByVideo   = "video.getAlbumsByVideo"
	MethodGetComments        = "video.getComments"
	MethodCreateComment      = "video.createComment"
	MethodDeleteComment      = "video.deleteComment"
	MethodRestoreComment     = "video.restoreComment"
	MethodEditComment        = "video.editComment"
	MethodGetTags            = "video.getTags"
	MethodPutTag             = "video.putTag"
	MethodRemoveTag          = "video.removeTag"
	MethodGetNewTags         = "video.getNewTags"
	MethodReport             = "video.report"
	MethodReportComment      = "video.reportComment"
	MethodGetCatalog         = "video.getCatalog"
	MethodGetCatalogSection  = "video.getCatalogSection"
	MethodHideCatalogSection = "video.hideCatalogSection"
)

type Service struct {
	Client *vk_api_go.Client
}

func New(Client *vk_api_go.Client) (s Service) {
	return Service{Client: Client}
}

func (s *Service) Get(r GetRequest) ([]byte, error) {
	return s.Client.Do(MethodGet, r)
}

func (s *Service) Edit(r EditRequest) ([]byte, error) {
	return s.Client.Do(MethodEdit, r)
}

func (s *Service) Add(r AddRequest) ([]byte, error) {
	return s.Client.Do(MethodAdd, r)
}

func (s *Service) Save(r SaveRequest) ([]byte, error) {
	return s.Client.Do(MethodSave, r)
}

func (s *Service) Delete(r DeleteRequest) ([]byte, error) {
	return s.Client.Do(MethodDelete, r)
}

func (s *Service) Restore(r RestoreRequest) ([]byte, error) {
	return s.Client.Do(MethodRestore, r)
}

func (s *Service) Search(r SearchRequest) ([]byte, error) {
	return s.Client.Do(MethodSearch, r)
}

func (s *Service) GetUserVideos(r GetUserVideosRequest) ([]byte, error) {
	return s.Client.Do(MethodGetUserVideos, r)
}

func (s *Service) GetAlbums(r GetAlbumsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetAlbums, r)
}

func (s *Service) GetAlbumById(r GetAlbumByIdRequest) ([]byte, error) {
	return s.Client.Do(MethodGetAlbumById, r)
}

func (s *Service) AddAlbum(r AddAlbumRequest) ([]byte, error) {
	return s.Client.Do(MethodAddAlbum, r)
}

func (s *Service) EditAlbum(r EditAlbumRequest) ([]byte, error) {
	return s.Client.Do(MethodEditAlbum, r)
}

func (s *Service) DeleteAlbum(r DeleteAlbumRequest) ([]byte, error) {
	return s.Client.Do(MethodDeleteAlbum, r)
}

func (s *Service) ReorderAlbums(r ReorderAlbumsRequest) ([]byte, error) {
	return s.Client.Do(MethodReorderAlbums, r)
}

func (s *Service) ReorderVideos(r ReorderVideosRequest) ([]byte, error) {
	return s.Client.Do(MethodReorderVideos, r)
}

func (s *Service) AddToAlbum(r AddToAlbumRequest) ([]byte, error) {
	return s.Client.Do(MethodAddToAlbum, r)
}

func (s *Service) RemoveFromAlbum(r RemoveFromAlbumRequest) ([]byte, error) {
	return s.Client.Do(MethodRemoveFromAlbum, r)
}

func (s *Service) GetAlbumsByVideo(r GetAlbumsByVideoRequest) ([]byte, error) {
	return s.Client.Do(MethodGetAlbumsByVideo, r)
}

func (s *Service) GetComments(r GetCommentsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetComments, r)
}

func (s *Service) CreateComment(r CreateCommentRequest) ([]byte, error) {
	return s.Client.Do(MethodCreateComment, r)
}

func (s *Service) DeleteComment(r DeleteCommentRequest) ([]byte, error) {
	return s.Client.Do(MethodDeleteComment, r)
}

func (s *Service) RestoreComment(r RestoreCommentRequest) ([]byte, error) {
	return s.Client.Do(MethodRestoreComment, r)
}

func (s *Service) EditComment(r EditCommentRequest) ([]byte, error) {
	return s.Client.Do(MethodEditComment, r)
}

func (s *Service) GetTags(r GetTagsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetTags, r)
}

func (s *Service) PutTag(r PutTagRequest) ([]byte, error) {
	return s.Client.Do(MethodPutTag, r)
}

func (s *Service) RemoveTag(r RemoveTagRequest) ([]byte, error) {
	return s.Client.Do(MethodRemoveTag, r)
}

func (s *Service) GetNewTags(r GetNewTagsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetNewTags, r)
}

func (s *Service) Report(r ReportRequest) ([]byte, error) {
	return s.Client.Do(MethodReport, r)
}

func (s *Service) ReportComment(r ReportCommentRequest) ([]byte, error) {
	return s.Client.Do(MethodReportComment, r)
}

func (s *Service) GetCatalog(r GetCatalogRequest) ([]byte, error) {
	return s.Client.Do(MethodGetCatalog, r)
}

func (s *Service) GetCatalogSection(r GetCatalogSectionRequest) ([]byte, error) {
	return s.Client.Do(MethodGetCatalogSection, r)
}

func (s *Service) HideCatalogSection(r HideCatalogSectionRequest) ([]byte, error) {
	return s.Client.Do(MethodHideCatalogSection, r)
}
