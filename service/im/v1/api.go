// Code generated by lark suite oapi sdk gen
package v1

import (
	"lost_found/api"
	"lost_found/api/core/request"
	"lost_found/api/core/response"
	"lost_found/core"
	"lost_found/core/config"
	"io"
)

type Service struct {
	conf              *config.Config
	Messages          *MessageService
	Chats             *ChatService
	ChatMemberUsers   *ChatMemberUserService
	ChatMemberBots    *ChatMemberBotService
	ChatAnnouncements *ChatAnnouncementService
	ChatMemberss      *ChatMembersService
	Files             *FileService
	Images            *ImageService
	MessageResources  *MessageResourceService
}

func NewService(conf *config.Config) *Service {
	s := &Service{
		conf: conf,
	}
	s.Messages = newMessageService(s)
	s.Chats = newChatService(s)
	s.ChatMemberUsers = newChatMemberUserService(s)
	s.ChatMemberBots = newChatMemberBotService(s)
	s.ChatAnnouncements = newChatAnnouncementService(s)
	s.ChatMemberss = newChatMembersService(s)
	s.Files = newFileService(s)
	s.Images = newImageService(s)
	s.MessageResources = newMessageResourceService(s)
	return s
}

type MessageService struct {
	service *Service
}

func newMessageService(service *Service) *MessageService {
	return &MessageService{
		service: service,
	}
}

type ChatService struct {
	service *Service
}

func newChatService(service *Service) *ChatService {
	return &ChatService{
		service: service,
	}
}

type ChatMemberUserService struct {
	service *Service
}

func newChatMemberUserService(service *Service) *ChatMemberUserService {
	return &ChatMemberUserService{
		service: service,
	}
}

type ChatMemberBotService struct {
	service *Service
}

func newChatMemberBotService(service *Service) *ChatMemberBotService {
	return &ChatMemberBotService{
		service: service,
	}
}

type ChatAnnouncementService struct {
	service *Service
}

func newChatAnnouncementService(service *Service) *ChatAnnouncementService {
	return &ChatAnnouncementService{
		service: service,
	}
}

type ChatMembersService struct {
	service *Service
}

func newChatMembersService(service *Service) *ChatMembersService {
	return &ChatMembersService{
		service: service,
	}
}

type FileService struct {
	service *Service
}

func newFileService(service *Service) *FileService {
	return &FileService{
		service: service,
	}
}

type ImageService struct {
	service *Service
}

func newImageService(service *Service) *ImageService {
	return &ImageService{
		service: service,
	}
}

type MessageResourceService struct {
	service *Service
}

func newMessageResourceService(service *Service) *MessageResourceService {
	return &MessageResourceService{
		service: service,
	}
}

type MessageListReqCall struct {
	ctx         *core.Context
	messages    *MessageService
	queryParams map[string]interface{}
	optFns      []request.OptFn
}

func (rc *MessageListReqCall) SetContainerIdType(containerIdType string) {
	rc.queryParams["container_id_type"] = containerIdType
}
func (rc *MessageListReqCall) SetContainerId(containerId string) {
	rc.queryParams["container_id"] = containerId
}
func (rc *MessageListReqCall) SetStartTime(startTime string) {
	rc.queryParams["start_time"] = startTime
}
func (rc *MessageListReqCall) SetEndTime(endTime string) {
	rc.queryParams["end_time"] = endTime
}
func (rc *MessageListReqCall) SetPageToken(pageToken string) {
	rc.queryParams["page_token"] = pageToken
}
func (rc *MessageListReqCall) SetPageSize(pageSize int) {
	rc.queryParams["page_size"] = pageSize
}

func (rc *MessageListReqCall) Do() (*MessageListResult, error) {
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &MessageListResult{}
	req := request.NewRequest("im/v1/messages", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.messages.service.conf, req)
	return result, err
}

func (messages *MessageService) List(ctx *core.Context, optFns ...request.OptFn) *MessageListReqCall {
	return &MessageListReqCall{
		ctx:         ctx,
		messages:    messages,
		queryParams: map[string]interface{}{},
		optFns:      optFns,
	}
}

type MessagePatchReqCall struct {
	ctx        *core.Context
	messages   *MessageService
	body       *MessagePatchReqBody
	pathParams map[string]interface{}
	optFns     []request.OptFn
}

func (rc *MessagePatchReqCall) SetMessageId(messageId string) {
	rc.pathParams["message_id"] = messageId
}

func (rc *MessagePatchReqCall) Do() (*response.NoData, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	var result = &response.NoData{}
	req := request.NewRequest("im/v1/messages/:message_id", "PATCH",
		[]request.AccessTokenType{request.AccessTokenTypeTenant, request.AccessTokenTypeUser}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.messages.service.conf, req)
	return result, err
}

func (messages *MessageService) Patch(ctx *core.Context, body *MessagePatchReqBody, optFns ...request.OptFn) *MessagePatchReqCall {
	return &MessagePatchReqCall{
		ctx:        ctx,
		messages:   messages,
		body:       body,
		pathParams: map[string]interface{}{},
		optFns:     optFns,
	}
}

type MessageReplyReqCall struct {
	ctx        *core.Context
	messages   *MessageService
	body       *MessageReplyReqBody
	pathParams map[string]interface{}
	optFns     []request.OptFn
}

func (rc *MessageReplyReqCall) SetMessageId(messageId string) {
	rc.pathParams["message_id"] = messageId
}

func (rc *MessageReplyReqCall) Do() (*Message, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	var result = &Message{}
	req := request.NewRequest("im/v1/messages/:message_id/reply", "POST",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.messages.service.conf, req)
	return result, err
}

func (messages *MessageService) Reply(ctx *core.Context, body *MessageReplyReqBody, optFns ...request.OptFn) *MessageReplyReqCall {
	return &MessageReplyReqCall{
		ctx:        ctx,
		messages:   messages,
		body:       body,
		pathParams: map[string]interface{}{},
		optFns:     optFns,
	}
}

type MessageCreateReqCall struct {
	ctx         *core.Context
	messages    *MessageService
	body        *MessageCreateReqBody
	queryParams map[string]interface{}
	optFns      []request.OptFn
}

func (rc *MessageCreateReqCall) SetReceiveIdType(receiveIdType string) {
	rc.queryParams["receive_id_type"] = receiveIdType
}

func (rc *MessageCreateReqCall) Do() (*Message, error) {
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &Message{}
	req := request.NewRequest("im/v1/messages", "POST",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.messages.service.conf, req)
	return result, err
}

func (messages *MessageService) Create(ctx *core.Context, body *MessageCreateReqBody, optFns ...request.OptFn) *MessageCreateReqCall {
	return &MessageCreateReqCall{
		ctx:         ctx,
		messages:    messages,
		body:        body,
		queryParams: map[string]interface{}{},
		optFns:      optFns,
	}
}

type MessageDeleteReqCall struct {
	ctx        *core.Context
	messages   *MessageService
	pathParams map[string]interface{}
	optFns     []request.OptFn
}

func (rc *MessageDeleteReqCall) SetMessageId(messageId string) {
	rc.pathParams["message_id"] = messageId
}

func (rc *MessageDeleteReqCall) Do() (*response.NoData, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	var result = &response.NoData{}
	req := request.NewRequest("im/v1/messages/:message_id", "DELETE",
		[]request.AccessTokenType{request.AccessTokenTypeTenant, request.AccessTokenTypeUser}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.messages.service.conf, req)
	return result, err
}

func (messages *MessageService) Delete(ctx *core.Context, optFns ...request.OptFn) *MessageDeleteReqCall {
	return &MessageDeleteReqCall{
		ctx:        ctx,
		messages:   messages,
		pathParams: map[string]interface{}{},
		optFns:     optFns,
	}
}

type MessageReadUsersReqCall struct {
	ctx         *core.Context
	messages    *MessageService
	pathParams  map[string]interface{}
	queryParams map[string]interface{}
	optFns      []request.OptFn
}

func (rc *MessageReadUsersReqCall) SetMessageId(messageId string) {
	rc.pathParams["message_id"] = messageId
}
func (rc *MessageReadUsersReqCall) SetUserIdType(userIdType string) {
	rc.queryParams["user_id_type"] = userIdType
}
func (rc *MessageReadUsersReqCall) SetPageSize(pageSize int) {
	rc.queryParams["page_size"] = pageSize
}
func (rc *MessageReadUsersReqCall) SetPageToken(pageToken string) {
	rc.queryParams["page_token"] = pageToken
}

func (rc *MessageReadUsersReqCall) Do() (*MessageReadUsersResult, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &MessageReadUsersResult{}
	req := request.NewRequest("im/v1/messages/:message_id/read_users", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.messages.service.conf, req)
	return result, err
}

func (messages *MessageService) ReadUsers(ctx *core.Context, optFns ...request.OptFn) *MessageReadUsersReqCall {
	return &MessageReadUsersReqCall{
		ctx:         ctx,
		messages:    messages,
		pathParams:  map[string]interface{}{},
		queryParams: map[string]interface{}{},
		optFns:      optFns,
	}
}

type ChatUpdateReqCall struct {
	ctx         *core.Context
	chats       *ChatService
	body        *ChatUpdateReqBody
	pathParams  map[string]interface{}
	queryParams map[string]interface{}
	optFns      []request.OptFn
}

func (rc *ChatUpdateReqCall) SetChatId(chatId string) {
	rc.pathParams["chat_id"] = chatId
}
func (rc *ChatUpdateReqCall) SetUserIdType(userIdType string) {
	rc.queryParams["user_id_type"] = userIdType
}

func (rc *ChatUpdateReqCall) Do() (*response.NoData, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &response.NoData{}
	req := request.NewRequest("im/v1/chats/:chat_id", "PUT",
		[]request.AccessTokenType{request.AccessTokenTypeTenant, request.AccessTokenTypeUser}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.chats.service.conf, req)
	return result, err
}

func (chats *ChatService) Update(ctx *core.Context, body *ChatUpdateReqBody, optFns ...request.OptFn) *ChatUpdateReqCall {
	return &ChatUpdateReqCall{
		ctx:         ctx,
		chats:       chats,
		body:        body,
		pathParams:  map[string]interface{}{},
		queryParams: map[string]interface{}{},
		optFns:      optFns,
	}
}

type FileCreateReqCall struct {
	ctx    *core.Context
	files  *FileService
	body   *request.FormData
	optFns []request.OptFn
}

func (rc *FileCreateReqCall) SetFileType(fileType string) {
	rc.body.AddParam("file_type", fileType)
}
func (rc *FileCreateReqCall) SetFileName(fileName string) {
	rc.body.AddParam("file_name", fileName)
}
func (rc *FileCreateReqCall) SetDuration(duration int) {
	rc.body.AddParam("duration", duration)
}
func (rc *FileCreateReqCall) SetFile(file *request.File) {
	rc.body.AddFile("file", file)
}

func (rc *FileCreateReqCall) Do() (*FileCreateResult, error) {
	var result = &FileCreateResult{}
	req := request.NewRequest("im/v1/files", "POST",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.files.service.conf, req)
	return result, err
}

func (files *FileService) Create(ctx *core.Context, optFns ...request.OptFn) *FileCreateReqCall {
	return &FileCreateReqCall{
		ctx:    ctx,
		files:  files,
		body:   request.NewFormData(),
		optFns: optFns,
	}
}

type FileGetReqCall struct {
	ctx        *core.Context
	files      *FileService
	pathParams map[string]interface{}
	optFns     []request.OptFn
	result     io.Writer
}

func (rc *FileGetReqCall) SetFileKey(fileKey string) {
	rc.pathParams["file_key"] = fileKey
}
func (rc *FileGetReqCall) SetResponseStream(result io.Writer) {
	rc.result = result
}

func (rc *FileGetReqCall) Do() (io.Writer, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	rc.optFns = append(rc.optFns, request.SetResponseStream())
	req := request.NewRequest("im/v1/files/:file_key", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, nil, rc.result, rc.optFns...)
	err := api.Send(rc.ctx, rc.files.service.conf, req)
	return rc.result, err
}

func (files *FileService) Get(ctx *core.Context, optFns ...request.OptFn) *FileGetReqCall {
	return &FileGetReqCall{
		ctx:        ctx,
		files:      files,
		pathParams: map[string]interface{}{},
		optFns:     optFns,
	}
}

type ChatListReqCall struct {
	ctx         *core.Context
	chats       *ChatService
	queryParams map[string]interface{}
	optFns      []request.OptFn
}

func (rc *ChatListReqCall) SetUserIdType(userIdType string) {
	rc.queryParams["user_id_type"] = userIdType
}
func (rc *ChatListReqCall) SetPageToken(pageToken string) {
	rc.queryParams["page_token"] = pageToken
}
func (rc *ChatListReqCall) SetPageSize(pageSize int) {
	rc.queryParams["page_size"] = pageSize
}

func (rc *ChatListReqCall) Do() (*ChatListResult, error) {
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &ChatListResult{}
	req := request.NewRequest("im/v1/chats", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeUser, request.AccessTokenTypeTenant}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.chats.service.conf, req)
	return result, err
}

func (chats *ChatService) List(ctx *core.Context, optFns ...request.OptFn) *ChatListReqCall {
	return &ChatListReqCall{
		ctx:         ctx,
		chats:       chats,
		queryParams: map[string]interface{}{},
		optFns:      optFns,
	}
}

type ImageCreateReqCall struct {
	ctx    *core.Context
	images *ImageService
	body   *request.FormData
	optFns []request.OptFn
}

func (rc *ImageCreateReqCall) SetImageType(imageType string) {
	rc.body.AddParam("image_type", imageType)
}
func (rc *ImageCreateReqCall) SetImage(image *request.File) {
	rc.body.AddFile("image", image)
}

func (rc *ImageCreateReqCall) Do() (*ImageCreateResult, error) {
	var result = &ImageCreateResult{}
	req := request.NewRequest("im/v1/images", "POST",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.images.service.conf, req)
	return result, err
}

func (images *ImageService) Create(ctx *core.Context, optFns ...request.OptFn) *ImageCreateReqCall {
	return &ImageCreateReqCall{
		ctx:    ctx,
		images: images,
		body:   request.NewFormData(),
		optFns: optFns,
	}
}

type ChatDeleteReqCall struct {
	ctx        *core.Context
	chats      *ChatService
	pathParams map[string]interface{}
	optFns     []request.OptFn
}

func (rc *ChatDeleteReqCall) SetChatId(chatId string) {
	rc.pathParams["chat_id"] = chatId
}

func (rc *ChatDeleteReqCall) Do() (*response.NoData, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	var result = &response.NoData{}
	req := request.NewRequest("im/v1/chats/:chat_id", "DELETE",
		[]request.AccessTokenType{request.AccessTokenTypeUser, request.AccessTokenTypeTenant}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.chats.service.conf, req)
	return result, err
}

func (chats *ChatService) Delete(ctx *core.Context, optFns ...request.OptFn) *ChatDeleteReqCall {
	return &ChatDeleteReqCall{
		ctx:        ctx,
		chats:      chats,
		pathParams: map[string]interface{}{},
		optFns:     optFns,
	}
}

type ImageGetReqCall struct {
	ctx        *core.Context
	images     *ImageService
	pathParams map[string]interface{}
	optFns     []request.OptFn
	result     io.Writer
}

func (rc *ImageGetReqCall) SetImageKey(imageKey string) {
	rc.pathParams["image_key"] = imageKey
}
func (rc *ImageGetReqCall) SetResponseStream(result io.Writer) {
	rc.result = result
}

func (rc *ImageGetReqCall) Do() (io.Writer, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	rc.optFns = append(rc.optFns, request.SetResponseStream())
	req := request.NewRequest("im/v1/images/:image_key", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, nil, rc.result, rc.optFns...)
	err := api.Send(rc.ctx, rc.images.service.conf, req)
	return rc.result, err
}

func (images *ImageService) Get(ctx *core.Context, optFns ...request.OptFn) *ImageGetReqCall {
	return &ImageGetReqCall{
		ctx:        ctx,
		images:     images,
		pathParams: map[string]interface{}{},
		optFns:     optFns,
	}
}

type ChatGetReqCall struct {
	ctx         *core.Context
	chats       *ChatService
	pathParams  map[string]interface{}
	queryParams map[string]interface{}
	optFns      []request.OptFn
}

func (rc *ChatGetReqCall) SetChatId(chatId string) {
	rc.pathParams["chat_id"] = chatId
}
func (rc *ChatGetReqCall) SetUserIdType(userIdType string) {
	rc.queryParams["user_id_type"] = userIdType
}

func (rc *ChatGetReqCall) Do() (*ChatGetResult, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &ChatGetResult{}
	req := request.NewRequest("im/v1/chats/:chat_id", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeUser, request.AccessTokenTypeTenant}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.chats.service.conf, req)
	return result, err
}

func (chats *ChatService) Get(ctx *core.Context, optFns ...request.OptFn) *ChatGetReqCall {
	return &ChatGetReqCall{
		ctx:         ctx,
		chats:       chats,
		pathParams:  map[string]interface{}{},
		queryParams: map[string]interface{}{},
		optFns:      optFns,
	}
}

type ChatCreateReqCall struct {
	ctx         *core.Context
	chats       *ChatService
	body        *ChatCreateReqBody
	queryParams map[string]interface{}
	optFns      []request.OptFn
}

func (rc *ChatCreateReqCall) SetUserIdType(userIdType string) {
	rc.queryParams["user_id_type"] = userIdType
}

func (rc *ChatCreateReqCall) Do() (*ChatCreateResult, error) {
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &ChatCreateResult{}
	req := request.NewRequest("im/v1/chats", "POST",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.chats.service.conf, req)
	return result, err
}

func (chats *ChatService) Create(ctx *core.Context, body *ChatCreateReqBody, optFns ...request.OptFn) *ChatCreateReqCall {
	return &ChatCreateReqCall{
		ctx:         ctx,
		chats:       chats,
		body:        body,
		queryParams: map[string]interface{}{},
		optFns:      optFns,
	}
}

type ChatSearchReqCall struct {
	ctx         *core.Context
	chats       *ChatService
	queryParams map[string]interface{}
	optFns      []request.OptFn
}

func (rc *ChatSearchReqCall) SetUserIdType(userIdType string) {
	rc.queryParams["user_id_type"] = userIdType
}
func (rc *ChatSearchReqCall) SetQuery(query string) {
	rc.queryParams["query"] = query
}
func (rc *ChatSearchReqCall) SetPageToken(pageToken string) {
	rc.queryParams["page_token"] = pageToken
}
func (rc *ChatSearchReqCall) SetPageSize(pageSize int) {
	rc.queryParams["page_size"] = pageSize
}

func (rc *ChatSearchReqCall) Do() (*ChatSearchResult, error) {
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &ChatSearchResult{}
	req := request.NewRequest("im/v1/chats/search", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeUser, request.AccessTokenTypeTenant}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.chats.service.conf, req)
	return result, err
}

func (chats *ChatService) Search(ctx *core.Context, optFns ...request.OptFn) *ChatSearchReqCall {
	return &ChatSearchReqCall{
		ctx:         ctx,
		chats:       chats,
		queryParams: map[string]interface{}{},
		optFns:      optFns,
	}
}

type MessageGetReqCall struct {
	ctx        *core.Context
	messages   *MessageService
	pathParams map[string]interface{}
	optFns     []request.OptFn
}

func (rc *MessageGetReqCall) SetMessageId(messageId string) {
	rc.pathParams["message_id"] = messageId
}

func (rc *MessageGetReqCall) Do() (*MessageGetResult, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	var result = &MessageGetResult{}
	req := request.NewRequest("im/v1/messages/:message_id", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.messages.service.conf, req)
	return result, err
}

func (messages *MessageService) Get(ctx *core.Context, optFns ...request.OptFn) *MessageGetReqCall {
	return &MessageGetReqCall{
		ctx:        ctx,
		messages:   messages,
		pathParams: map[string]interface{}{},
		optFns:     optFns,
	}
}

type ChatMembersCreateReqCall struct {
	ctx          *core.Context
	chatMemberss *ChatMembersService
	body         *ChatMembersCreateReqBody
	pathParams   map[string]interface{}
	queryParams  map[string]interface{}
	optFns       []request.OptFn
}

func (rc *ChatMembersCreateReqCall) SetChatId(chatId string) {
	rc.pathParams["chat_id"] = chatId
}
func (rc *ChatMembersCreateReqCall) SetMemberIdType(memberIdType string) {
	rc.queryParams["member_id_type"] = memberIdType
}

func (rc *ChatMembersCreateReqCall) Do() (*ChatMembersCreateResult, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &ChatMembersCreateResult{}
	req := request.NewRequest("im/v1/chats/:chat_id/members", "POST",
		[]request.AccessTokenType{request.AccessTokenTypeUser, request.AccessTokenTypeTenant}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.chatMemberss.service.conf, req)
	return result, err
}

func (chatMemberss *ChatMembersService) Create(ctx *core.Context, body *ChatMembersCreateReqBody, optFns ...request.OptFn) *ChatMembersCreateReqCall {
	return &ChatMembersCreateReqCall{
		ctx:          ctx,
		chatMemberss: chatMemberss,
		body:         body,
		pathParams:   map[string]interface{}{},
		queryParams:  map[string]interface{}{},
		optFns:       optFns,
	}
}

type ChatMembersDeleteReqCall struct {
	ctx          *core.Context
	chatMemberss *ChatMembersService
	body         *ChatMembersDeleteReqBody
	pathParams   map[string]interface{}
	queryParams  map[string]interface{}
	optFns       []request.OptFn
}

func (rc *ChatMembersDeleteReqCall) SetChatId(chatId string) {
	rc.pathParams["chat_id"] = chatId
}
func (rc *ChatMembersDeleteReqCall) SetMemberIdType(memberIdType string) {
	rc.queryParams["member_id_type"] = memberIdType
}

func (rc *ChatMembersDeleteReqCall) Do() (*ChatMembersDeleteResult, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &ChatMembersDeleteResult{}
	req := request.NewRequest("im/v1/chats/:chat_id/members", "DELETE",
		[]request.AccessTokenType{request.AccessTokenTypeUser, request.AccessTokenTypeTenant}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.chatMemberss.service.conf, req)
	return result, err
}

func (chatMemberss *ChatMembersService) Delete(ctx *core.Context, body *ChatMembersDeleteReqBody, optFns ...request.OptFn) *ChatMembersDeleteReqCall {
	return &ChatMembersDeleteReqCall{
		ctx:          ctx,
		chatMemberss: chatMemberss,
		body:         body,
		pathParams:   map[string]interface{}{},
		queryParams:  map[string]interface{}{},
		optFns:       optFns,
	}
}

type ChatMembersGetReqCall struct {
	ctx          *core.Context
	chatMemberss *ChatMembersService
	pathParams   map[string]interface{}
	queryParams  map[string]interface{}
	optFns       []request.OptFn
}

func (rc *ChatMembersGetReqCall) SetChatId(chatId string) {
	rc.pathParams["chat_id"] = chatId
}
func (rc *ChatMembersGetReqCall) SetMemberIdType(memberIdType string) {
	rc.queryParams["member_id_type"] = memberIdType
}
func (rc *ChatMembersGetReqCall) SetPageToken(pageToken string) {
	rc.queryParams["page_token"] = pageToken
}
func (rc *ChatMembersGetReqCall) SetPageSize(pageSize int) {
	rc.queryParams["page_size"] = pageSize
}

func (rc *ChatMembersGetReqCall) Do() (*ChatMembersGetResult, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &ChatMembersGetResult{}
	req := request.NewRequest("im/v1/chats/:chat_id/members", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeUser, request.AccessTokenTypeTenant}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.chatMemberss.service.conf, req)
	return result, err
}

func (chatMemberss *ChatMembersService) Get(ctx *core.Context, optFns ...request.OptFn) *ChatMembersGetReqCall {
	return &ChatMembersGetReqCall{
		ctx:          ctx,
		chatMemberss: chatMemberss,
		pathParams:   map[string]interface{}{},
		queryParams:  map[string]interface{}{},
		optFns:       optFns,
	}
}

type ChatAnnouncementGetReqCall struct {
	ctx               *core.Context
	chatAnnouncements *ChatAnnouncementService
	pathParams        map[string]interface{}
	queryParams       map[string]interface{}
	optFns            []request.OptFn
}

func (rc *ChatAnnouncementGetReqCall) SetChatId(chatId string) {
	rc.pathParams["chat_id"] = chatId
}
func (rc *ChatAnnouncementGetReqCall) SetUserIdType(userIdType string) {
	rc.queryParams["user_id_type"] = userIdType
}

func (rc *ChatAnnouncementGetReqCall) Do() (*ChatAnnouncementGetResult, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &ChatAnnouncementGetResult{}
	req := request.NewRequest("im/v1/chats/:chat_id/announcement", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeUser, request.AccessTokenTypeTenant}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.chatAnnouncements.service.conf, req)
	return result, err
}

func (chatAnnouncements *ChatAnnouncementService) Get(ctx *core.Context, optFns ...request.OptFn) *ChatAnnouncementGetReqCall {
	return &ChatAnnouncementGetReqCall{
		ctx:               ctx,
		chatAnnouncements: chatAnnouncements,
		pathParams:        map[string]interface{}{},
		queryParams:       map[string]interface{}{},
		optFns:            optFns,
	}
}

type MessageResourceGetReqCall struct {
	ctx              *core.Context
	messageResources *MessageResourceService
	pathParams       map[string]interface{}
	queryParams      map[string]interface{}
	optFns           []request.OptFn
	result           io.Writer
}

func (rc *MessageResourceGetReqCall) SetMessageId(messageId string) {
	rc.pathParams["message_id"] = messageId
}
func (rc *MessageResourceGetReqCall) SetFileKey(fileKey string) {
	rc.pathParams["file_key"] = fileKey
}
func (rc *MessageResourceGetReqCall) SetType(type_ string) {
	rc.queryParams["type"] = type_
}
func (rc *MessageResourceGetReqCall) SetResponseStream(result io.Writer) {
	rc.result = result
}

func (rc *MessageResourceGetReqCall) Do() (io.Writer, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	rc.optFns = append(rc.optFns, request.SetResponseStream())
	req := request.NewRequest("im/v1/messages/:message_id/resources/:file_key", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, nil, rc.result, rc.optFns...)
	err := api.Send(rc.ctx, rc.messageResources.service.conf, req)
	return rc.result, err
}

func (messageResources *MessageResourceService) Get(ctx *core.Context, optFns ...request.OptFn) *MessageResourceGetReqCall {
	return &MessageResourceGetReqCall{
		ctx:              ctx,
		messageResources: messageResources,
		pathParams:       map[string]interface{}{},
		queryParams:      map[string]interface{}{},
		optFns:           optFns,
	}
}

type ChatMembersIsInChatReqCall struct {
	ctx          *core.Context
	chatMemberss *ChatMembersService
	pathParams   map[string]interface{}
	optFns       []request.OptFn
}

func (rc *ChatMembersIsInChatReqCall) SetChatId(chatId string) {
	rc.pathParams["chat_id"] = chatId
}

func (rc *ChatMembersIsInChatReqCall) Do() (*ChatMembersIsInChatResult, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	var result = &ChatMembersIsInChatResult{}
	req := request.NewRequest("im/v1/chats/:chat_id/members/is_in_chat", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeUser, request.AccessTokenTypeTenant}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.chatMemberss.service.conf, req)
	return result, err
}

func (chatMemberss *ChatMembersService) IsInChat(ctx *core.Context, optFns ...request.OptFn) *ChatMembersIsInChatReqCall {
	return &ChatMembersIsInChatReqCall{
		ctx:          ctx,
		chatMemberss: chatMemberss,
		pathParams:   map[string]interface{}{},
		optFns:       optFns,
	}
}

type ChatMembersMeJoinReqCall struct {
	ctx          *core.Context
	chatMemberss *ChatMembersService
	pathParams   map[string]interface{}
	optFns       []request.OptFn
}

func (rc *ChatMembersMeJoinReqCall) SetChatId(chatId string) {
	rc.pathParams["chat_id"] = chatId
}

func (rc *ChatMembersMeJoinReqCall) Do() (*response.NoData, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	var result = &response.NoData{}
	req := request.NewRequest("im/v1/chats/:chat_id/members/me_join", "PATCH",
		[]request.AccessTokenType{request.AccessTokenTypeUser, request.AccessTokenTypeTenant}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.chatMemberss.service.conf, req)
	return result, err
}

func (chatMemberss *ChatMembersService) MeJoin(ctx *core.Context, optFns ...request.OptFn) *ChatMembersMeJoinReqCall {
	return &ChatMembersMeJoinReqCall{
		ctx:          ctx,
		chatMemberss: chatMemberss,
		pathParams:   map[string]interface{}{},
		optFns:       optFns,
	}
}

type ChatAnnouncementPatchReqCall struct {
	ctx               *core.Context
	chatAnnouncements *ChatAnnouncementService
	body              *ChatAnnouncementPatchReqBody
	pathParams        map[string]interface{}
	optFns            []request.OptFn
}

func (rc *ChatAnnouncementPatchReqCall) SetChatId(chatId string) {
	rc.pathParams["chat_id"] = chatId
}

func (rc *ChatAnnouncementPatchReqCall) Do() (*response.NoData, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	var result = &response.NoData{}
	req := request.NewRequest("im/v1/chats/:chat_id/announcement", "PATCH",
		[]request.AccessTokenType{request.AccessTokenTypeUser, request.AccessTokenTypeTenant}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.chatAnnouncements.service.conf, req)
	return result, err
}

func (chatAnnouncements *ChatAnnouncementService) Patch(ctx *core.Context, body *ChatAnnouncementPatchReqBody, optFns ...request.OptFn) *ChatAnnouncementPatchReqCall {
	return &ChatAnnouncementPatchReqCall{
		ctx:               ctx,
		chatAnnouncements: chatAnnouncements,
		body:              body,
		pathParams:        map[string]interface{}{},
		optFns:            optFns,
	}
}
