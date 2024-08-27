import axios from 'axios';

// URL del server che gestisce il login
const API_URL = 'http://localhost:5173/session';

export const login = async (username) => {
  try {
    const response = await axios.post(API_URL, { username });
    return response.data; // Restituisce i dati della risposta (es. token)
  } catch (error) {
    // Stampa dettagli aggiuntivi sull'errore
    console.error('Errore durante la richiesta di login:', error.message);
    if (error.response) {
      console.error('Dettagli risposta errore:', error.response.data);
      console.error('Status codice:', error.response.status);
    } else if (error.request) {
      console.error('Richiesta non ricevuta:', error.request);
    } else {
      console.error('Errore:', error.message);
    }
    throw new Error(error.response?.data?.message || 'Errore di login');
  }
};