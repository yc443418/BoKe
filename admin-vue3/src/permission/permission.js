/**
 * 获取权限配置
 */
import { checkAuthority } from "@/utils/authority"
import storage from "../utils/storage"

export const hasPermission = {
    install: (app) => {
        app.directive('hasPermission', {
            mounted(el, binding) {
                const value = binding.value
                const permissions = storage.getItem('permissionList') || []
                const hasPermission = checkAuthority(value, permissions)
                if (!hasPermission) {
                    if (el.parentNode) {
                        el.parentNode.removeChild(el)
                    } else {
                        el.innerHTML = ''
                    }
                } else {
                    el && el.setAttribute('code', value)
                }
            }
        })
    }
}