import {createRouter, createWebHashHistory} from 'vue-router'
import SessionView from '../views/SessionView.vue'
import LoginView from '../views/LoginView.vue'
import ProfileView from '../views/ProfileView.vue'



const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: LoginView},
		{path: '/session', component: SessionView},
		{path: '/users/:username/profile', component: ProfileView},
		
	]
})

export default router
