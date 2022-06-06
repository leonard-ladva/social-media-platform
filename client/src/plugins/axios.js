import axios from 'axios'

axios.interceptors.request.use(
	config => {
		const token = localStorage.getItem('token');
		const auth = token ? `Bearer ${token}` : '';
		config.headers.common['Authorization'] = auth;
		return config;
	},
	error => Promise.reject(error),
);

axios.defaults.baseURL = 'http://localhost:9100/';

export default axios