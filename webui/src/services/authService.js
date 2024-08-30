// src/services/authservice.js
import axios from 'axios';

const API_URL = 'http://localhost:3000/session'; // Assicurati che questo sia il percorso corretto

export async function doLogin(username) {
  console.log("ciao");
  try {
    console.log("!")
    const response = await axios.post(API_URL,{ username }, // Corpo della richiesta
      { headers: { 'Authorization': `Bearer ${sessionStorage.getItem('token')}` } } // Headers della richiesta
    );
    console.log(response.data)
    console.log("!")
    return response.data;
    
  } catch (error) {
    // Gestione dettagliata degli errori
    let errorMessage = 'Errore di login: ';
    if (error.response) {
      // La richiesta è stata fatta e il server ha risposto con uno stato
      // che non è nel range di 2xx
      errorMessage += error.response.data.message || error.response.statusText;
    } else if (error.request) {
      // La richiesta è stata fatta ma non è stata ricevuta alcuna risposta
      errorMessage += 'Nessuna risposta dal server.';
    } else {
      // Qualcosa è andato storto nella configurazione della richiesta
      errorMessage += error.message;
    }
    throw new Error(errorMessage);
  }
}
