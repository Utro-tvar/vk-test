import React, { useEffect, useState } from 'react';
import axios from 'axios';

function Table({backendUrl, period}) {
    const [data, setData] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    const fetchData = async () => {
        try {
            setLoading(true);
            const response = await axios.get(`${backendUrl}/read`);
            setData(response.data);
        } catch (error) {
            setError(`Error while fetching data from ${backendUrl}/read`);
        }finally{
            setLoading(false);
        }
    };

    useEffect(() => {

        fetchData();
        const interval = setInterval(fetchData, period);

        return () => clearInterval(interval);
    }, []);


    if (loading) {
        return <div className="text-primary" role="status"><span className="sr-only">Loading...</span></div>
    }
    if (error) {
        return <div className="alert alert-danger" role="alert">Error: {error}</div>;
    }
    return (
        <div className="container mt-4">
            <h2>containers</h2>
            <table className="table table-bordered table-striped">
                <thead>
                    <tr>
                        <th>IP</th>
                        <th>Ping</th>
                        <th>Last connection</th>
                    </tr>
                </thead>
                <tbody>
                    {data.map(cont => (
                        <tr key={cont.ip}>
                            <td>{cont.ip}</td>
                            <td>{cont.ping}</td>
                            <td>{cont.last_conn}</td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
}
export default Table;