import React from 'react';
import moment from 'moment';

const TransactionList = ({ transactions = [], onDelete, onSelect }) => {
      const formatCurrency = (amount) =>
            new Intl.NumberFormat('id-ID', {
                  style: 'currency',
                  currency: 'IDR',
                  minimumFractionDigits: 0,
                  maximumFractionDigits: 0
            }).format(amount);

      const formatDate = (date) => moment(date).format('DD MMM YYYY, HH:mm');

      return (
            <div className="bg-white shadow rounded-lg p-6 mt-6">
                  <h2 className="text-xl font-semibold mb-4">Transaction History</h2>
                  {(transactions || []).length === 0 ? (
                        <p className="text-center text-gray-500">No transactions found.</p>
                  ) : (
                        <div className="overflow-x-auto">
                              <table className="min-w-full table-auto border border-gray-200">
                                    <thead>
                                          <tr className="bg-gray-100">
                                                <th className="text-left px-4 py-2 border">No</th>
                                                <th className="text-left px-4 py-2 border">Date</th>
                                                <th className="text-left px-4 py-2 border">Type</th>
                                                <th className="text-left px-4 py-2 border">Category</th>
                                                <th className="text-left px-4 py-2 border">Amount</th>
                                                <th className="text-left px-4 py-2 border">Description</th>
                                                <th className="text-left px-4 py-2 border">Actions</th>
                                          </tr>
                                    </thead>
                                    <tbody>
                                          {transactions.map((transaction) => (
                                                <tr
                                                      key={transaction.id}
                                                      className="hover:bg-gray-50 cursor-pointer"
                                                      onClick={() => onSelect(transaction)}
                                                >
                                                      <td className="px-4 py-2 border">{transactions.indexOf(transaction) + 1}</td>
                                                      <td className="px-4 py-2 border">{formatDate(transaction.date)}</td>
                                                      <td className="px-4 py-2 border">
                                                            <span
                                                                  className={
                                                                        transaction.type === 'income'
                                                                              ? 'text-green-600 font-medium'
                                                                              : 'text-red-600 font-medium'
                                                                  }
                                                            >
                                                                  {transaction.type.charAt(0).toUpperCase() + transaction.type.slice(1)}
                                                            </span>
                                                      </td>
                                                      <td className="px-4 py-2 border">{transaction.category}</td>
                                                      <td
                                                            className={`px-4 py-2 border ${transaction.type === 'income' ? 'text-green-600' : 'text-red-600'
                                                                  }`}
                                                      >
                                                            {formatCurrency(transaction.amount)}
                                                      </td>
                                                      <td className="px-4 py-2 border">{transaction.description || '-'}</td>
                                                      <td className="px-4 py-2 border">
                                                            <button
                                                                  onClick={(e) => {
                                                                        e.stopPropagation();
                                                                        onDelete(transaction.id);
                                                                  }}
                                                                  className="bg-red-500 text-white px-3 py-1 rounded hover:bg-red-600 text-sm"
                                                            >
                                                                  Delete
                                                            </button>
                                                      </td>
                                                </tr>
                                          ))}
                                    </tbody>
                              </table>
                        </div>
                  )}
            </div>
      );
};

export default TransactionList;
