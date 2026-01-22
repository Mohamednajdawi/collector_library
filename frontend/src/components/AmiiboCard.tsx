import React from 'react';
import { Amiibo } from '../types';

interface AmiiboCardProps {
    amiibo: Amiibo;
}

const AmiiboCard: React.FC<AmiiboCardProps> = ({ amiibo }) => {
    return (
        <div className="amiibo-card">
            <div className="image-container">
                <img
                    src={amiibo.image_url}
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
