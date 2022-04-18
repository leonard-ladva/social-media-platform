import axios from 'axios'

axios.defaults.baseURL = 'http://localhost:9100/';
axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token');

export default axios