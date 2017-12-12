package photos

import (
	"github.com/drossha/vk-api-go"
)

const (
	MethodSave           = "photos.save"
	MethodCopy           = "photos.copy"
	MethodEdit           = "photos.edit"
	MethodMove           = "photos.move"
	MethodMakeCover      = "photos.makeCover"
	MethodReorderAlbums  = "photos.reorderAlbums"
	MethodReorderPhotos  = "photos.reorderPhotos"
	MethodGetAll         = "photos.getAll"
	MethodGetUserPhotos  = "photos.getUserPhotos"
	MethodDeleteAlbum    = "photos.deleteAlbum"
	MethodDelete         = "photos.delete"
	MethodRestore        = "photos.restore"
	MethodConfirmTag     = "photos.confirmTag"
	MethodGetComments    = "photos.getComments"
	MethodGetAllComments = "photos.getAllComments"
	MethodCreateComment  = "photos.createComment"
	MethodDeleteComment  = "photos.deleteComment"
	MethodRestoreComment = "photos.restoreComment"
	MethodEditComment    = "photos.editComment"
	MethodGetTags        = "photos.getTags"
	MethodPutTag         = "photos.putTag"
	MethodRemoveTag      = "photos.removeTag"
	MethodGetNewTags     = "photos.getNewTags"
)

type Service struct {
	Client *vk_api_go.Client
}

func New(Client *vk_api_go.Client) (s Service) {
	return Service{Client: Client}
}

func (s *Service) Save(r SaveRequest) ([]byte, error) {
	return s.Client.Do(MethodSave, r)
}

func (s *Service) Copy(r CopyRequest) ([]byte, error) {
	return s.Client.Do(MethodCopy, r)
}

func (s *Service) Edit(r EditRequest) ([]byte, error) {
	return s.Client.Do(MethodEdit, r)
}

func (s *Service) Move(r MoveRequest) ([]byte, error) {
	return s.Client.Do(MethodMove, r)
}

func (s *Service) MakeCover(r MakeCoverRequest) ([]byte, error) {
	return s.Client.Do(MethodMakeCover, r)
}

func (s *Service) ReorderAlbums(r ReorderAlbumsRequest) ([]byte, error) {
	return s.Client.Do(MethodReorderAlbums, r)
}

func (s *Service) ReorderPhotos(r ReorderPhotosRequest) ([]byte, error) {
	return s.Client.Do(MethodReorderPhotos, r)
}

func (s *Service) GetAll(r GetAllRequest) ([]byte, error) {
	return s.Client.Do(MethodGetAll, r)
}

func (s *Service) GetUserPhotos(r GetUserPhotosRequest) ([]byte, error) {
	return s.Client.Do(MethodGetUserPhotos, r)
}

func (s *Service) DeleteAlbum(r DeleteAlbumRequest) ([]byte, error) {
	return s.Client.Do(MethodDeleteAlbum, r)
}

func (s *Service) Delete(r DeleteRequest) ([]byte, error) {
	return s.Client.Do(MethodDelete, r)
}

func (s *Service) Restore(r RestoreRequest) ([]byte, error) {
	return s.Client.Do(MethodRestore, r)
}

func (s *Service) ConfirmTag(r ConfirmTagRequest) ([]byte, error) {
	return s.Client.Do(MethodConfirmTag, r)
}

func (s *Service) GetComments(r GetCommentsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetComments, r)
}

func (s *Service) GetAllComments(r GetAllCommentsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetAllComments, r)
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
