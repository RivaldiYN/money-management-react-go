import React, { useState, useEffect } from 'react';

const TransactionForm = ({ onAdd, onUpdate, transaction }) => {
      const [formData, setFormData] = useState({
            type: 'income',
            amount: '',
            category: '',
            description: ''
      });

      const [isEditing, setIsEditing] = useState(false);

      useEffect(() => {
            if (transaction) {
                  setFormData({
                        type: transaction.type,
                        amount: transaction.amount,
                        category: transaction.category,
                        description: transaction.description || ''
                  });
                  setIsEditing(true);
            } else {
                  resetForm();
            }
      }, [transaction]);

      const handleChange = (e) => {
            const { name, value } = e.target;
            setFormData({
                  ...formData,
                  [name]: name === 'amount' ? parseFloat(value) || '' : value
            });
      };

      const handleSubmit = (e) => {
            e.preventDefault();
            if (!formData.amount || !formData.category) return;

            if (isEditing && transaction) {
                  onUpdate(transaction.id, formData);
            } else {
                  onAdd(formData);
            }

            resetForm();
      };

      const resetForm = () => {
            setFormData({
                  type: 'income',
                  amount: '',
                  category: '',
                  description: ''
            });
            setIsEditing(false);
      };

      return (
            <div className="bg-white shadow rounded-lg p-6">
                  <h2 className="text-xl font-semibold mb-4">
                        {isEditing ? 'Edit Transaction' : 'Add New Transaction'}
                  </h2>
                  <form onSubmit={handleSubmit} className="space-y-4">
                        <div>
                              <label className="block font-medium mb-1">Type</label>
                              <select
                                    name="type"
                                    value={formData.type}
                                    onChange={handleChange}
                                    className="w-full border rounded px-3 py-2"
                                    required
                              >
                                    <option value="income">Income</option>
                                    <option value="expense">Expense</option>
                              </select>
                        </div>

                        <div>
                              <label className="block font-medium mb-1">Amount</label>
                              <input
                                    type="number"
                                    name="amount"
                                    value={formData.amount}
                                    onChange={handleChange}
                                    placeholder="Enter amount"
                                    min="0"
                                    step="0.01"
                                    className="w-full border rounded px-3 py-2"
                                    required
                              />
                        </div>

                        <div>
                              <label className="block font-medium mb-1">Category</label>
                              <input
                                    type="text"
                                    name="category"
                                    value={formData.category}
                                    onChange={handleChange}
                                    placeholder="Enter category"
                                    className="w-full border rounded px-3 py-2"
                                    required
                              />
                        </div>

                        <div>
                              <label className="block font-medium mb-1">Description</label>
                              <textarea
                                    name="description"
                                    value={formData.description}
                                    onChange={handleChange}
                                    placeholder="Enter description (optional)"
                                    rows={3}
                                    className="w-full border rounded px-3 py-2"
                              />
                        </div>

                        <div className="space-y-2">
                              <button
                                    type="submit"
                                    className="w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700 transition"
                              >
                                    {isEditing ? 'Update Transaction' : 'Add Transaction'}
                              </button>

                              {isEditing && (
                                    <button
                                          type="button"
                                          onClick={resetForm}
                                          className="w-full bg-gray-500 text-white py-2 rounded hover:bg-gray-600 transition"
                                    >
                                          Cancel Editing
                                    </button>
                              )}
                        </div>
                  </form>
            </div>
      );
};

export default TransactionForm;
