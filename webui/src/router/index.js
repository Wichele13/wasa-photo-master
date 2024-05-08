import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginVue from '../views/LoginVue.vue'
import ProfileVue from '../views/ProfileVue.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: LoginVue},
		{path: '/users/:username/profile', component: ProfileVue},
		
	]
})

export default router
