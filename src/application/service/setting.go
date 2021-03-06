package service

import (
	"context"
	"errors"
	"github.com/Etpmls/EM-Auth/src/application/client"
	"github.com/Etpmls/EM-Auth/src/application/protobuf"
	em "github.com/Etpmls/Etpmls-Micro"
	em_library "github.com/Etpmls/Etpmls-Micro/library"
	em_protobuf "github.com/Etpmls/Etpmls-Micro/protobuf"

	"google.golang.org/grpc/codes"
)

type ServiceSetting struct {
	protobuf.UnimplementedSettingServer
}

// Clear all cache
// 清除全部缓存
func (this *ServiceSetting) CacheClear(ctx context.Context, request *em_protobuf.Empty) (*em_protobuf.Response, error) {
	// If the cache is not turned on, return to the prompt
	// 如果没开启缓存，返回提示
	if !em_library.Config.App.Cache {
		return em.ErrorRpc(codes.InvalidArgument, em.ERROR_Code, em_library.I18n.TranslateFromRequest(ctx, "ERROR_MESSAGE_CacheIsNotEnabled"), nil, errors.New("Cache Is Not Enabled!"))
	}

	em_library.Cache.ClearAllCache()
	em.LogDebug.Output(em.MessageWithLineNum("Cleared all cache!"))

	return em.SuccessRpc(em.SUCCESS_Code, em_library.I18n.TranslateFromRequest(ctx, "SUCCESS_Clear"), nil)
}

// Disk Cleanup
// 清理磁盘
func (this *ServiceSetting) DiskCleanUp(ctx context.Context, request *em_protobuf.Empty) (*em_protobuf.Response, error) {
	err := client.NewClient().Setting_DiskCleanUp(ctx)
	if err != nil {
		return em.ErrorRpc(codes.InvalidArgument, em.ERROR_Code, em_library.I18n.TranslateFromRequest(ctx, "ERROR_Delete"), nil, err)
	}

	em.LogDebug.Output(em.MessageWithLineNum("Disk cleanup complete!"))
	return em.SuccessRpc(em.SUCCESS_Code, em_library.I18n.TranslateFromRequest(ctx, "SUCCESS_Clear"), nil)
}