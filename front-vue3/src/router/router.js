import {
    createRouter,
    createWebHistory
} from 'vue-router'

/**
 * 路由表
 */
const router = createRouter({
    history: createWebHistory(),
    routes: [{
        path: '',
        component: () => import('@/views/layout.vue'),
        redirect: 'home',
        children: [{
                path: '/home',
                component: () => import('@/views/home.vue'),
            },
            {
                path: '/category',
                component: () => import('@/views/category.vue'),
            },
            {
                path: '/tag',
                component: () => import('@/views/tag.vue'),
            },
            {
                path: '/create/article',
                component: () => import('@/views/createArticle.vue'),
            },
            {
                path: '/search',
                component: () => import('@/views/search.vue'),
            },
            {
                path: '/user/:userId',
                component: ()=> import('@/views/user.vue'),
            },
            {
                path: '/personal',
                component: ()=> import('@/views/personal.vue'),
            },
            {
                path: '/article/:articleId',
                component: ()=> import('@/views/articleDetail.vue'),
            },
            {
                path: '/edit/article/:articleId',
                component: ()=> import('@/views/editArticle.vue'),
            },
            {
                path: '/user/message/:type',
                component: ()=> import('@/views/message.vue'),
            },
        ]
    }]
})

export default router