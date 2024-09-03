<template>
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
      <li class="list-group-item" v-for="user in users" :key="user.id">
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
      error: null    // Stato di errore
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
    }
  }
};
</script>

<style scoped>
/* Aggiunta dello stile */
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
