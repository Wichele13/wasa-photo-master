import {createRouter, createWebHashHistory} from 'vue-router'
import SessionVue from '../views/SessionVue.vue'
import LoginVue from '../views/LoginVue.vue'
import UserProfileVue from '../views/UserProfileVue.vue'
import SearchProfilesVue from '../views/SearchProfilesVue.vue'
import ErroreVue from '../views/ErroreVue.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: LoginVue},
		{path: '/session', component: SessionVue},
		{path: '/users/:username/profile', component: UserProfileVue},
		{path: '/users/:username/view', component: SearchProfilesVue},
		{path: '/:catchAll(.*)', component: ErroreVue}
	]
})

export default router
