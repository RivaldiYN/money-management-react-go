import React, { useState, useEffect } from 'react';
import Dashboard from './components/Dashboard';
import TransactionForm from './components/TransactionForm';
import TransactionList from './components/TransactionList';
import api from './services/api';

function App() {
  const [transactions, setTransactions] = useState([]);
  const [summary, setSummary] = useState({ totalIncome: 0, totalExpense: 0, balance: 0 });
  const [error, setError] = useState('');
  const [showAlert, setShowAlert] = useState(false);
  const [alertMessage, setAlertMessage] = useState('');
  const [selectedTransaction, setSelectedTransaction] = useState(null);

  useEffect(() => {
    const fetchInitialData = async () => {
      try {
        const transactionsRes = await api.getTransactions();
        setTransactions(transactionsRes.data.data || []);

        const summaryRes = await api.getSummary();
        const { total_income, total_expense, balance } = summaryRes.data.data;
        setSummary({
          totalIncome: total_income,
          totalExpense: total_expense,
          balance: balance,
        });
      } catch (err) {
        setError('Gagal mengambil data awal');
      }
    };

    fetchInitialData();
  }, []);


  const fetchTransactions = async () => {
    try {
      const response = await api.getTransactions();
      setTransactions(response.data.data || []);
    } catch (err) {
      setError('Error fetching transactions');
      setTransactions([]);
    }
  };

  const fetchSummary = async () => {
    try {
      const response = await api.getSummary();
      setSummary(response.data.data);
    } catch (err) {
      setError('Error fetching summary');
    }
  };

  const handleAddTransaction = async (transaction) => {
    try {
      await api.createTransaction(transaction);
      setAlertMessage('Transaction added successfully');
      setShowAlert(true);
      fetchTransactions();
      fetchSummary();
      setTimeout(() => setShowAlert(false), 3000);
    } catch (err) {
      setError('Error adding transaction');
    }
  };

  const handleUpdateTransaction = async (id, transaction) => {
    try {
      await api.updateTransaction(id, transaction);
      setAlertMessage('Transaction updated successfully');
      setShowAlert(true);
      fetchTransactions();
      fetchSummary();
      setSelectedTransaction(null);
      setTimeout(() => setShowAlert(false), 3000);
    } catch (err) {
      setError('Error updating transaction');
    }
  };

  const handleDeleteTransaction = async (id) => {
    try {
      await api.deleteTransaction(id);
      setAlertMessage('Transaction deleted successfully');
      setShowAlert(true);
      fetchTransactions();
      fetchSummary();
      setTimeout(() => setShowAlert(false), 3000);
    } catch (err) {
      setError('Error deleting transaction');
    }
  };

  const handleSelectTransaction = (transaction) => {
    setSelectedTransaction(transaction);
  };

  return (
    <div className="min-h-screen bg-gray-100 p-6">
      <h1 className="text-3xl font-bold text-center text-gray-800 mb-6">Money Management App</h1>

      {showAlert && (
        <div className="max-w-xl mx-auto mb-4 px-4 py-3 bg-green-100 text-green-700 rounded">
          {alertMessage}
        </div>
      )}

      {error && (
        <div className="max-w-xl mx-auto mb-4 px-4 py-3 bg-red-100 text-red-700 rounded">
          {error}
        </div>
      )}

      <div className="mb-6">
        <Dashboard summary={summary} />
      </div>

      <div className="grid md:grid-cols-3 gap-6">
        <div>
          <TransactionForm
            onAdd={handleAddTransaction}
            onUpdate={handleUpdateTransaction}
            transaction={selectedTransaction}
          />
        </div>
        <div className="md:col-span-2">
          <TransactionList
            transactions={transactions}
            onDelete={handleDeleteTransaction}
            onSelect={handleSelectTransaction}
          />
        </div>
      </div>
    </div>
  );
}

export default App;
