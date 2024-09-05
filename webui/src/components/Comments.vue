<script>
export default {
	props: ["log", "token"],
	data() {
		return {
		}
	},

	methods: {
		async deleteComment(username, photoid, commentid) {
			try {
				let response = await this.$axios.delete("/users/" + username + "/photos/" + photoid + "/comments/" + commentid, {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("token")
					}
				})
				location.reload();
			} catch (e) {
				if (e.response && e.response.status === 400) {
					this.errormsg = "Form error, please check all fields and try again. If you think that this is an error, write an e-mail to us.";
					this.detailedmsg = null;
				} else if (e.response && e.response.status === 500) {
					this.errormsg = "An internal error occurred. We will be notified. Please try again later.";
					this.detailedmsg = e.toString();
				} else {
					this.errormsg = e.toString();
					this.detailedmsg = null;
				}
			}
		},
	},
	mounted() {
	}
}
</script>

<template>
	<div class="modal fade" tabindex="-1">
		<div class="modal-dialog modal-xl modal-dialog-centered">
			<div class="modal-content">
				<div class="modal-header bg-dark text-white">
					<h5 class="modal-title">Comments</h5>
					<button type="button" class="btn-close btn-close-white" data-bs-dismiss="modal" aria-label="Close"></button>
				</div>
				<div class="modal-body">
					<div class="row g-4">
						<div class="col-md-4" v-for="comment in log.comments" :key="comment.id">
							<div class="card h-100 shadow-sm border-0">
								<div class="card-body d-flex flex-column">
									<h5 class="card-title">{{ comment.username }}</h5>
									<p class="card-text flex-grow-1">{{ comment.content }}</p>
									<button v-if="token == comment.userId" type="button" class="btn btn-danger mt-3 align-self-end" @click="deleteComment(comment.ownerUsername, comment.photoId, comment.id)">Delete</button>
								</div>
							</div>
						</div>
					</div>
				</div>
				<div class="modal-footer">
					<button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
				</div>
			</div>
		</div>
	</div>
</template>

<style scoped>
.modal-header {
	background-color: #343a40;
	color: #fff;
}
.modal-body {
	background-color: #f8f9fa;
}
.card {
	background-color: #ffffff;
	border-radius: 0.75rem;
}
.card-title {
	font-size: 1.25rem;
	font-weight: bold;
	color: #343a40;
}
.card-text {
	color: #6c757d;
}
.btn-danger {
	background-color: #dc3545;
	border-color: #dc3545;
}
.btn-secondary {
	background-color: #6c757d;
	border-color: #6c757d;
}
.btn-close-white {
	filter: invert(1);
}
</style>
