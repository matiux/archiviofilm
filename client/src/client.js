import axios from "axios";

const client = axios.create({
   baseURL: 'http://localhost:8080/api/v1',
   timeout: 1000,
   //responseType: 'json'
});

export default client;

export function fetchList() {

   return client.get('/film');
}

export function updateFilm(filmId, props) {

   return client.patch('/film/' + filmId, props);
}