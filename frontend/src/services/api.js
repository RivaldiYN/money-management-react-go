import axios from 'axios';

const API_URL = 'http://localhost:8080/api/v1';

const api = {
      getTransactions: () => axios.get(`${API_URL}/transactions`),
      getTransaction: (id) => axios.get(`${API_URL}/transactions/${id}`),
      createTransaction: (data) => axios.post(`${API_URL}/transactions`, data),
      updateTransaction: (id, data) => axios.put(`${API_URL}/transactions/${id}`, data),
      deleteTransaction: (id) => axios.delete(`${API_URL}/transactions/${id}`),
      getSummary: () => axios.get(`${API_URL}/summary`),
};

export default api;
