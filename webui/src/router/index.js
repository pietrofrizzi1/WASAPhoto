import {createRouter, createWebHashHistory} from 'vue-router'
import SessionView from '../views/SessionView.vue'
import ErrorView from '../views/ErrorView.vue'
import LoginView from '../views/LoginView.vue'



const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		
		
		{path: '/:catchAll(.*)', component: ErrorView},
		{path: '/', name: 'login',component: LoginView,},
		{
			path: '/session',
			name: 'Session',
			component: SessionView,
		  },
		  {
			path: '/',
			redirect: '/login', // Reindirizza la radice alla pagina di login
		  },
        
		
	]
})

export default router
