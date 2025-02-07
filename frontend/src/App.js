import React, { useEffect, useState } from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import Table from './components/Table';

function App() {
  const [error, setError] = useState(null);
  const [loading, setLoading] = useState(true);

  let backendUrl = "test";
  let period = 5000;

  useEffect(() => {
    backendUrl = process.env.BACKEND_URL;
    if (!backendUrl) {
      setError('Backend URL is not defined!');
    }

    period = process.env.PERIOD;
    if (!period) {
      setError('Period is not defined!');
    }
    setLoading(false)
  }, []);

  if (loading) {
    return <div className="text-primary" role="status"><span className="sr-only">Loading...</span></div>
  }
  if (error) {
    return <div className="alert alert-danger" role="alert">Error: {error}</div>;
  }

  return (
    <div className="container">
      <h1 className="my-4">Ping info about docker containers</h1>
      <Table backendUrl={backendUrl} period={period} />
    </div>
  );
}

export default App;