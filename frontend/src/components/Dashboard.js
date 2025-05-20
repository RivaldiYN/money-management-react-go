const Dashboard = ({ summary }) => {
      // Format currency
      const formatCurrency = (amount) => {
            return new Intl.NumberFormat('id-ID', {
                  style: 'currency',
                  currency: 'IDR',
                  minimumFractionDigits: 0,
                  maximumFractionDigits: 0
            }).format(amount);
      };

      return (
            <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
                  <div className="bg-white shadow rounded-lg p-6">
                        <h2 className="text-lg font-semibold text-gray-700">Total Income</h2>
                        <p className="mt-2 text-2xl font-bold text-green-500">{isNaN(summary.totalIncome) ? 'Loading...' : formatCurrency(summary.totalIncome)}</p>
                  </div>
                  <div className="bg-white shadow rounded-lg p-6">
                        <h2 className="text-lg font-semibold text-gray-700">Total Expense</h2>
                        <p className="mt-2 text-2xl font-bold text-red-500">{isNaN(summary.totalExpense) ? 'Loading...' : formatCurrency(summary.totalExpense)}</p>
                  </div>
                  <div className="bg-white shadow rounded-lg p-6">
                        <h2 className="text-lg font-semibold text-gray-700">Balance</h2>
                        <p className="mt-2 text-2xl font-bold text-blue-500">{isNaN(summary.balance) ? 'Loading...' : formatCurrency(summary.balance)}</p>
                  </div>
            </div>
      );
};
export default Dashboard;