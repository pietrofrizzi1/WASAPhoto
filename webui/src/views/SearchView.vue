
<template>
 <button @click="toggleSidebar" class="hamburger-btn">
    <svg width="30" height="30" viewBox="0 0 30 30">
        <rect width="30" height="5"></rect>
        <rect y="12" width="30" height="5"></rect>
        <rect y="24" width="30" height="5"></rect>
    </svg>
 </button>
  <nav id="sidebarMenu" class="sidebar bg-light" v-if="isSidebarOpen">
    <div class="position-sticky pt-3 sidebar-sticky">
        <h6 class="sidebar-heading text-muted text-uppercase">General</h6>
        <ul class="nav flex-column">
            <li class="nav-item">
                <RouterLink to="/session" class="nav-link">
                    <svg class="feather">
                        <use href="/feather-sprite-v4.29.0.svg#home" />
                    </svg>
                    Home
                </RouterLink>
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


  <div class="search-container">
    <h1>Risultati della ricerca</h1>
    <p v-if="query">Mostrando i risultati per: {{ query }}</p>
    
    <!-- Mostra un messaggio di caricamento durante la ricerca -->
    <div v-if="loading" class="spinner-border" role="status">
      <span class="sr-only">Caricamento...</span>
    </div>
    
    <!-- Mostra un messaggio di errore se c'è un problema -->
    <div v-if="error" class="alert alert-danger">{{ error }}</div>
    
    <!-- Mostra i risultati della ricerca -->
    <ul class="list-group" v-if="users.length">
      <li class="list-group-item" v-for="user in users" :key="user.id" @click="loaduserprofile(user.username)">
        {{ user.username }}
      </li>
    </ul>
    
    <!-- Mostra un messaggio se non ci sono risultati -->
    <p v-if="!loading && !users.length && query">Nessun risultato trovato per "{{ query }}"</p>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  props: ['query'],
  data() {
    return {
      username: localStorage.getItem('username'),
			token: localStorage.getItem('token'),
      users: [],    // Array per memorizzare i risultati della ricerca
      loading: false,  // Stato di caricamento
      error: null,
      isSidebarOpen: false, 
    };
  },
  watch: {
    query(newQuery) {
      // Chiama la funzione di ricerca ogni volta che la query cambia
      if (newQuery) {
        this.searchUsers(newQuery);
      }
    }
  },
  mounted() {
    // Avvia la ricerca se una query è presente al montaggio del componente
    if (this.query) {
      this.searchUsers(this.query);
    }
  },
  methods: {
    async searchUsers(username) {
      this.loading = true;
      this.error = null;
      try {
        // Esegui una richiesta GET al server per cercare gli utenti
        let response = await this.$axios.get('/users/'+this.username+'/search/'+this.query, {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("token")
					}
				})
			
        this.users = response.data;  // Assume che la risposta contenga un array di utenti
      } catch (err) {
        this.error = 'Errore durante il recupero dei dati di ricerca';
      } finally {
        this.loading = false;
      }
    },
    async loaduserprofile(user) {
      this.$router.push({ path: '/users/' + user + '/view' })
    
    },
    toggleSidebar() {
      this.isSidebarOpen = !this.isSidebarOpen;
    },
    async doLogout() {
			localStorage.removeItem("token")
			localStorage.removeItem("username")
			this.$router.push({ path: '/' })
		},
  }
};
</script>

<style scoped>
.hamburger-btn {
  background: none;
  border: none;
  cursor: pointer;
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  width: 30px;
  height: 30px;
  padding: 0;
}

.hamburger-btn span {
  display: block;
  width: 100%;
  height: 3px;
  background-color: black; 
  border-radius: 2px;
  transition: background-color 0.3s;
}

.search-container {
  margin-top: 20px;
  max-width: 800px;
  margin-left: auto;
  margin-right: auto;
}

.spinner-border {
  width: 3rem;
  height: 3rem;
}

.alert {
  margin-top: 20px;
}

.list-group-item {
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.list-group-item:hover {
  background-color: #f8f9fa;
}
</style>
