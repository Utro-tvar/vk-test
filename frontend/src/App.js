import React, { useEffect, useState } from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import Table from './components/Table';

function App() {
  const [error, setError] = useState(null);
  const [loading, setLoading] = useState(true);

  let backendUrl = "test";
  let period = 5000;

  useEffect(() => {
    console.log(process.env)

    backendUrl = process.env.REACT_APP_BACKEND_URL;
    if (!backendUrl) {
      setError('Backend URL is not defined!');
    }

    const periodStr = process.env.REACT_APP_PERIOD;
    if (!periodStr) {
      setError('Period is not defined!');
    }
    period = Number(periodStr)
    if(isNaN(period)){
      setError(`${periodStr}`)
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