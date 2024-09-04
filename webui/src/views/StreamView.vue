<script>
import Comments from "../components/Comments.vue";
import SuccessMsg from "../components/SuccessMsg.vue";

export default {
	components: { Comments, SuccessMsg },
	data: function () {
		return {
			errormsg: null,
			successmsg: null,
			detailedmsg: null,
			username: localStorage.getItem('username'),
			token: localStorage.getItem('token'),
			loading: false,
			some_data: null,
			images: null,
			image: null,
			clear: null,
			photoComments: {
				requestIdentifier: 0,
				photoIdentifier: 0,
				identifier: 0,
				comments: [
					{
						id: 0,
						userId: 0,
						photoId: 0,
						photoOwner: 0,
						ownerUsername: "",
						username: "",
						content: "",
					}
				],
			},
			comment: "",
			stream: {
				identifier: 0,
				photoStream: [
					{
						id: 0,
						userId: 0,
						username: "",
						file: "",
						date: "",
						likeCount: 0,
						commentCount: 0,
						comment: "",
						likeStatus: null,
					}
				],
			},
			searchUserUsername: "",
			like: {
				likeId: 0,
				identifier: 0,
				photoIdentifier: 0,
				photoOwner: 0,
			},
			profile: {
				requestId: 0,
				id: 0,
				username: "",
				followersCount: 0,
				followingCount: 0,
				photoCount: 0,
				followStatus: null,
				banStatus: null,
			},
		}
	},
	methods: {
		async refresh() {
			this.getStream()
		},
		async getStream() {
			try {
				let response = await this.$axios.get("/users/" + this.username + "/stream", {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("token")
					}
				})
				this.stream = response.data
				for (let i = 0; i < this.stream.photoStream.length; i++) {
					this.stream.photoStream[i].file = 'data:image/*;base64,' + this.stream.photoStream[i].file
				}
			} catch (e) {
				if (e.response && e.response.status === 400) {
					this.errormsg = "Form error, please check all fields and try again. If you think that this is an error, write an e-mail to us.";
					this.detailedmsg = null;
				} else if (e.response && e.response.status === 500) {
					this.errormsg = "An internal error occurred. We will be notified. Please try again later.";
					this.detailedmsg = e.toString();
				} else {
					this.errormsg = "You are not following anyone yet. Follow someone to see their photos.";
					this.detailedmsg = null;
				}
			}
		},
		async sendComment(username, photoid, comment) {
			if (comment === "") {
				this.errormsg = "Emtpy comment field."
			} else {
				try {
					let response = await this.$axios.put("/users/" + username + "/photos/" + photoid + "/comments/" + Math.floor(Math.random() * 10000), { content: comment }, {
						headers: {
							Authorization: "Bearer " + localStorage.getItem("token")
						}
					})
					this.clear = response.data
					this.refresh()
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
			}
		},
		async openLog(username, photoid) {
			try {
				let response = await this.$axios.get("/users/" + username + "/photos/" + photoid + "/comments", {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("token")
					}
				})
				this.photoComments = response.data;
				const modal = new bootstrap.Modal(document.getElementById('logviewer'));
				modal.show();
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
		async likePhoto(username, id) {
			try {
				let response = await this.$axios.put("/users/" + username + "/photos/" + id + "/likes/" + Math.floor(Math.random() * 10000), {}, {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("token")
					}
				})
				this.clear = response.data
				this.refresh()
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
		async deleteLike(username, id) {
			try {
				let response = await this.$axios.get("/users/" + username + "/photos/" + id + "/likes", {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("token")
					}
				})
				this.like = response.data
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

			try {
				let response = await this.$axios.delete("/users/" + username + "/photos/" + id + "/likes/" + this.like.likeId, {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("token")
					}
				})
				this.clear = response.data
				this.refresh()
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
		async doLogout() {
			localStorage.removeItem("token")
			localStorage.removeItem("username")
			this.$router.push({ path: '/' })
		},
		async ViewProfile() {
			this.$router.push({ path: '/users/' + this.username + '/profile' })
		},
		async SearchUser(){
			
			if (this.searchUserUsername === "") {

				this.errormsg = "Emtpy username field."
			} else {
			   
			   
			   this.$router.push({ path: '/users/'+this.username+'/search', query: { q: this.searchUserUsername } })

			}
		}
	},
	mounted() {
		this.getStream()
	}
}
</script>

<template>
	<div class="container">
	  <!-- Sidebar -->
	  
	  
	
	  <nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
		
		<div class="position-sticky pt-3 sidebar-sticky">
		  <ul class="nav flex-column">
			<li class="nav-item">
            <div class="nav-link" @click="ViewProfile">
              <svg class="feather">
                <use href="/feather-sprite-v4.29.0.svg#user" />
              </svg>
              Profile
            </div>
          </li>
          <li class="nav-item">
            <div class="nav-link" @click="doLogout">
              <svg class="feather">
                <use href="/feather-sprite-v4.29.0.svg#log-out" />
              </svg>
              Logout
            </div>
          </li>
        </ul>
      </div>
    </nav>
	  <!-- Main Content -->
	  <div class="col-md-9 ms-sm-auto col-lg-10 px-4">
		<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
		  <h1 class="h2">{{ username }}'s stream</h1>
		  <div class="btn-toolbar mb-2 mb-md-0">
		  </div>
		</div>
  
		<!-- Search User -->
		<div class="input-group mb-3">
		  <input type="text" v-model="searchUserUsername" class="form-control" placeholder="Search a user in WASAPhoto">
		  <button class="btn btn-primary" type="button" @click="SearchUser">Search</button>
		</div>
  
		<!-- Error and Success Messages -->
		<ErrorMsg v-if="errormsg" :msg="errormsg" class="mb-3"></ErrorMsg>
		<SuccessMsg v-if="successmsg" :msg="successmsg" class="mb-3"></SuccessMsg>
  
		<!-- Photo Stream -->
		<div class="row">
		  <div class="col-md-4 mb-4" v-for="photo in stream.photoStream" :key="photo.id">
			<div class="card shadow-sm">
			  <img class="card-img-top" :src="photo.file" alt="Photo" />
			  <div class="card-body">
				<RouterLink :to="'/users/' + photo.username + '/view'" class="nav-link">
				  <button type="button" class="btn btn-outline-primary">{{ photo.username }}</button>
				</RouterLink>
				<p class="card-text">Likes: {{ photo.likeCount }}</p>
				<p class="card-text">Comments: {{ photo.commentCount }}</p>
				<p class="card-text">Uploaded on: {{ photo.date }}</p>
  
				<!-- Comment Input -->
				<div class="input-group mb-3">
				  <input type="text" v-model="photo.comment" class="form-control" placeholder="Add a comment">
				  <button class="btn btn-primary" type="button" @click="sendComment(photo.username, photo.id, photo.comment)">Send</button>
				</div>
  
				<!-- Actions -->
				<div class="d-flex justify-content-between align-items-center">
				  <div class="btn-group">
					<button type="button" class="btn btn-dark" @click="openLog(photo.username, photo.id)">Comments</button>
					<button type="button" v-if="photo.likeStatus == false" class="btn btn-primary" @click="likePhoto(photo.username, photo.id)">Like</button>
					<button type="button" v-if="photo.likeStatus == true" class="btn btn-danger" @click="deleteLike(photo.username, photo.id)">Unlike</button>
				  </div>
				</div>
			  </div>
			</div>
		  </div>
		</div>
	  </div>
  
	  <!-- Modal -->
	  <Comments id="logviewer" :log="photoComments" :token="token"></Comments>
	</div>
</template>

<style scoped>
/* Container adjustments */
.container {
	margin-top: 20px;
}

/* Card and image styling */
.card {
	border-radius: 10px;
	overflow: hidden;
}
.card-img-top {
	object-fit: cover;
	height: 200px;
}

/* Input and button styling */
.input-group input {
	border-radius: 5px 0 0 5px;
}
.input-group button {
	border-radius: 0 5px 5px 0;
}

/* Button adjustments */
.btn-group .btn {
	margin-right: 5px;
}

/* Error and success messages */
.ErrorMsg, .SuccessMsg {
	border: 1px solid transparent;
	border-radius: 5px;
	padding: 10px;
}
.ErrorMsg {
	background-color: #f8d7da;
	color: #721c24;
}
.SuccessMsg {
	background-color: #d4edda;
	color: #155724;
}


@media (max-width: 767px) {
	.col-md-9 {
	  padding: 0;
	}
}

/* Sidebar adjustments */
.sidebar .nav-link {
	display: flex;
	align-items: center;
	padding: 10px 15px;
	border-radius: 5px;
}
.sidebar .nav-link svg {
	width: 20px;
	height: 20px;
	margin-right: 10px;
}
.sidebar .nav-link:hover {
	background-color: #e9ecef;
}
</style>
