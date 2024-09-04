import {createRouter, createWebHashHistory} from 'vue-router'
import StreamView from '../views/StreamView.vue'
import LoginView from '../views/LoginView.vue'
import MyProfileView from '../views/MyProfileView.vue'
import ProfilesView from '../views/ProfilesView.vue'
import ErorrView from '../views/404view.vue'
import SearchView from '../views/SearchView.vue';

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: LoginView},
		{path: '/session', component: StreamView},
		{path: '/users/:username/profile', component: MyProfileView},
		{path: '/users/:username/search', component: SearchView, props: route => ({ query: route.query.q }) },
		{path: '/users/:username/view', component: ProfilesView},
		{path: '/:catchAll(.*)', component: ErorrView}
	]
})

export default router
