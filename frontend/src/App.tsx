import { useEffect, useState } from 'react';
import Grid from './components/Grid';
import { Amiibo } from './types';

function App() {
    const [amiibos, setAmiibos] = useState<Amiibo[]>([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const fetchAmiibos = async () => {
            try {
                // In development, might typically point to localhost:8080
                // For now, let's assume the user will proxy or run on parallel ports
                // Use Runtime Config (injected at startup by Docker) or fallback to build-time/local
                // @ts-ignore
                const apiUrl = window.runtimeEnv?.API_URL || import.meta.env.VITE_API_URL || 'http://localhost:8080/api/amiibos';
                const response = await fetch(apiUrl);
                if (!response.ok) {
                    throw new Error('Failed to fetch');
                }
                const data = await response.json();
                setAmiibos(data);
            } catch (error) {
                console.error('Error fetching amiibos:', error);
            } finally {
                setLoading(false);
            }
        };

        fetchAmiibos();
    }, []);

    return (
        <div className="app">
            <header className="header">
                <div className="logo">COLLECTOR LIBRARY</div>
                {/* Placeholder for future nav or search */}
            </header>

            <main>
                {loading ? (
                    <div className="loading-container">Loading Collection...</div>
                ) : (
                    <Grid amiibos={amiibos} />
                )}
            </main>
        </div>
    );
}

export default App;
