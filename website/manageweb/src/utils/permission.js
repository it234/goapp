import store from '@/store'

/**
 * @param {Array} value
 * @returns {Boolean}
 * @example see @/views/permission/directive.vue
 */
export default function checkPermission(value) {
  if (value && value instanceof Array && value.length > 0) {
    const roles = store.getters && store.getters.roles
    const permissionRoles = value

    const hasPermission = roles.some(role => {
      return permissionRoles.includes(role)
    })

    if (!hasPermission) {
      return false
    }
    return true
  } else {
    console.error(`need roles! Like v-permission="['admin','editor']"`)
    return false
  }
}

/**
 * @param {Array} arr ['add','del','view','update']
 * @param {String} value 'add'
 * @returns {Boolean}
 * @example see @/views/permission/directive.vue
 */
export function checkAuth(arr, value) {
  const permissionarr = arr
  const permissionvalue = value
  const hasPermission = permissionarr.includes(permissionvalue)
  if (!hasPermission) {
    return false
  }
  return true
}
export function checkAuthAdd(arr) {
  return checkAuth(arr, 'add')
}
export function checkAuthDel(arr) {
  return checkAuth(arr, 'del')
}
export function checkAuthView(arr) {
  return checkAuth(arr, 'view')
}
export function checkAuthUpdate(arr) {
  return checkAuth(arr, 'update')
}
export function checkAuthSetadminrole(arr) {
  return checkAuth(arr, 'setadminrole')
}
export function checkAuthSetrolemenu(arr) {
  return checkAuth(arr, 'setrolemenu')
}
