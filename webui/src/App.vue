<script setup>
import { RouterLink, RouterView } from 'vue-router'

</script>
<script>
export default {
	data() {
		return {
		};

	},
}
</script>

<template>
	<div>
		<header class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
			<a class="navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-6" href="#/">WASAPhoto</a>
			
		</header>

		<div class="container-fluid">
			<div class="row">
				<main >
					<RouterView />
				</main>
			</div>
		</div>

	</div>

</template>

<style>

</style>
