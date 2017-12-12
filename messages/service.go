package messages

import (
	"github.com/drossha/vk-api-go"
)

const (
	MethodGet                        = "messages.get"
	MethodGetDialogs                 = "messages.getDialogs"
	MethodGetById                    = "messages.getById"
	MethodSearch                     = "messages.search"
	MethodGetHistory                 = "messages.getHistory"
	MethodGetHistoryAttachments      = "messages.getHistoryAttachments"
	MethodSend                       = "messages.send"
	MethodDelete                     = "messages.delete"
	MethodDeleteDialog               = "messages.deleteDialog"
	MethodRestore                    = "messages.restore"
	MethodMarkAsRead                 = "messages.markAsRead"
	MethodMarkAsImportant            = "messages.markAsImportant"
	MethodMarkAsImportantDialog      = "messages.markAsImportantDialog"
	MethodMarkAsUnansweredDialog     = "messages.markAsUnansweredDialog"
	MethodGetLongPollServer          = "messages.getLongPollServer"
	MethodGetLongPollHistory         = "messages.getLongPollHistory"
	MethodGetChat                    = "messages.getChat"
	MethodCreateChat                 = "messages.createChat"
	MethodEditChat                   = "messages.editChat"
	MethodGetChatUsers               = "messages.getChatUsers"
	MethodSetActivity                = "messages.setActivity"
	MethodSearchDialogs              = "messages.searchDialogs"
	MethodAddChatUser                = "messages.addChatUser"
	MethodRemoveChatUser             = "messages.removeChatUser"
	MethodGetLastActivity            = "messages.getLastActivity"
	MethodSetChatPhoto               = "messages.setChatPhoto"
	MethodDeleteChatPhoto            = "messages.deleteChatPhoto"
	MethodDenyMessagesFromGroup      = "messages.denyMessagesFromGroup"
	MethodAllowMessagesFromGroup     = "messages.allowMessagesFromGroup"
	MethodIsMessagesFromGroupAllowed = "messages.isMessagesFromGroupAllowed"
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

func (s *Service) GetDialogs(r GetDialogsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetDialogs, r)
}

func (s *Service) GetById(r GetByIdRequest) ([]byte, error) {
	return s.Client.Do(MethodGetById, r)
}

func (s *Service) Search(r SearchRequest) ([]byte, error) {
	return s.Client.Do(MethodSearch, r)
}

func (s *Service) GetHistory(r GetHistoryRequest) ([]byte, error) {
	return s.Client.Do(MethodGetHistory, r)
}

func (s *Service) GetHistoryAttachments(r GetHistoryAttachmentsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetHistoryAttachments, r)
}

func (s *Service) Send(r SendRequest) ([]byte, error) {
	return s.Client.Do(MethodSend, r)
}

func (s *Service) Delete(r DeleteRequest) ([]byte, error) {
	return s.Client.Do(MethodDelete, r)
}

func (s *Service) DeleteDialog(r DeleteDialogRequest) ([]byte, error) {
	return s.Client.Do(MethodDeleteDialog, r)
}

func (s *Service) Restore(r RestoreRequest) ([]byte, error) {
	return s.Client.Do(MethodRestore, r)
}

func (s *Service) MarkAsRead(r MarkAsReadRequest) ([]byte, error) {
	return s.Client.Do(MethodMarkAsRead, r)
}

func (s *Service) MarkAsImportant(r MarkAsImportantRequest) ([]byte, error) {
	return s.Client.Do(MethodMarkAsImportant, r)
}

func (s *Service) MarkAsImportantDialog(r MarkAsImportantDialogRequest) ([]byte, error) {
	return s.Client.Do(MethodMarkAsImportantDialog, r)
}

func (s *Service) MarkAsUnansweredDialog(r MarkAsUnansweredDialogRequest) ([]byte, error) {
	return s.Client.Do(MethodMarkAsUnansweredDialog, r)
}

func (s *Service) GetLongPollServer(r GetLongPollServerRequest) ([]byte, error) {
	return s.Client.Do(MethodGetLongPollServer, r)
}

func (s *Service) GetLongPollHistory(r GetLongPollHistoryRequest) ([]byte, error) {
	return s.Client.Do(MethodGetLongPollHistory, r)
}

func (s *Service) GetChat(r GetChatRequest) ([]byte, error) {
	return s.Client.Do(MethodGetChat, r)
}

func (s *Service) CreateChat(r CreateChatRequest) ([]byte, error) {
	return s.Client.Do(MethodCreateChat, r)
}

func (s *Service) EditChat(r EditChatRequest) ([]byte, error) {
	return s.Client.Do(MethodEditChat, r)
}

func (s *Service) GetChatUsers(r GetChatUsersRequest) ([]byte, error) {
	return s.Client.Do(MethodGetChatUsers, r)
}

func (s *Service) SetActivity(r SetActivityRequest) ([]byte, error) {
	return s.Client.Do(MethodSetActivity, r)
}

func (s *Service) SearchDialogs(r SearchDialogsRequest) ([]byte, error) {
	return s.Client.Do(MethodSearchDialogs, r)
}

func (s *Service) AddChatUser(r AddChatUserRequest) ([]byte, error) {
	return s.Client.Do(MethodAddChatUser, r)
}

func (s *Service) RemoveChatUser(r RemoveChatUserRequest) ([]byte, error) {
	return s.Client.Do(MethodRemoveChatUser, r)
}

func (s *Service) GetLastActivity(r GetLastActivityRequest) ([]byte, error) {
	return s.Client.Do(MethodGetLastActivity, r)
}

func (s *Service) SetChatPhoto(r SetChatPhotoRequest) ([]byte, error) {
	return s.Client.Do(MethodSetChatPhoto, r)
}

func (s *Service) DeleteChatPhoto(r DeleteChatPhotoRequest) ([]byte, error) {
	return s.Client.Do(MethodDeleteChatPhoto, r)
}

func (s *Service) DenyMessagesFromGroup(r DenyMessagesFromGroupRequest) ([]byte, error) {
	return s.Client.Do(MethodDenyMessagesFromGroup, r)
}

func (s *Service) AllowMessagesFromGroup(r AllowMessagesFromGroupRequest) ([]byte, error) {
	return s.Client.Do(MethodAllowMessagesFromGroup, r)
}

func (s *Service) IsMessagesFromGroupAllowed(r IsMessagesFromGroupAllowedRequest) ([]byte, error) {
	return s.Client.Do(MethodIsMessagesFromGroupAllowed, r)
}
