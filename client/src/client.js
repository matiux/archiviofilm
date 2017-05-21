import axios from "axios";

const client = axios.create({
   baseURL: 'http://localhost:8080/api/v1',
   timeout: 1000,
   //responseType: 'json'
});

export default client;

export function fetchList(unseen, filter) {

   const query = {
      params: {}
   };

   if (unseen) {

      query.params.unseen = true;
   }

   if (filter) {

      query.params.filters = filter;
   }

   return client.get('/film', query);
}

export function updateFilm(filmId, props) {

   return client.patch('/film/' + filmId, props);
}