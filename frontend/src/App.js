import React, { useState } from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import Table from './components/Table';

function App() {
  const [error, setError] = useState(null);

  const backendUrl = process.env.BACKEND_URL;
  if (!backendUrl) {
    setError('Backend URL is not defined!');
  }

  const period = process.env.PERIOD;
  if (!period) {
    setError('Period is not defined!');
  }

  return (
    <div className="container">
      <h1 className="my-4">Ping info about docker containers</h1>

      {error ? (
        <div className="alert alert-danger">{error}</div>
      ) : (
        <Table backendUrl={backendUrl} period={period} />
      )}
    </div>
  );
}

export default App;