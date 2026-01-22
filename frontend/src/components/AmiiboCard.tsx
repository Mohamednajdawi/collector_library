import React from 'react';
import { Amiibo } from '../types';

interface AmiiboCardProps {
    amiibo: Amiibo;
}

const AmiiboCard: React.FC<AmiiboCardProps> = ({ amiibo }) => {
    // Helper to fix Wikia/Fandom hotlinking issues by requesting a specific thumbnail version
    const getImageUrl = (url: string) => {
        if (url.includes('static.wikia.nocookie.net') && !url.includes('/revision/latest')) {
            return `${url}/revision/latest/scale-to-width-down/350?path-prefix=en`;
        }
        return url;
    };

    return (
        <div className="amiibo-card">
            <div className="image-container">
                <img
                    src={getImageUrl(amiibo.image_url)}
                    alt={amiibo.name}
                    className="amiibo-image"
                    loading="lazy"
                />
            </div>
            <div className="card-content">
                <div className="card-series">{amiibo.series}</div>
                <h3 className="card-title">{amiibo.name}</h3>
            </div>
        </div>
    );
};

export default AmiiboCard;
