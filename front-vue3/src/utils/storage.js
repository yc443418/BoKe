/**
 * 本地存储方式（存储，获取，删除方法）
*/
export default {
    getStorage() {
        return JSON.parse(window.localStorage.getItem(import.meta.env.VITE_APP_NAME_SPACE) || "{}")
    },
    setItem(key,val) {
        let storage = this.getStorage()
        storage[key] = val
        window.localStorage.setItem(import.meta.env.VITE_APP_NAME_SPACE, JSON.stringify(storage))
    },
    getItem(key) {
        return this.getStorage()[key]
    },
    clearItem(key) {
        let storage = this.getStorage()
        delete storage[key]
        window.localStorage.setItem(import.meta.env.VITE_APP_NAME_SPACE, JSON.stringify(storage))
    },
    clearAll(){
        window.localStorage.clear()
    }
}