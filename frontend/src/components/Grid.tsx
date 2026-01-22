import React from 'react';
import AmiiboCard from './AmiiboCard';
import { Amiibo } from '../types';

interface GridProps {
    amiibos: Amiibo[];
}

const Grid: React.FC<GridProps> = ({ amiibos }) => {
    return (
        <div className="grid-container">
            {amiibos.map((amiibo) => (
                <AmiiboCard key={amiibo.id || amiibo.name} amiibo={amiibo} />
            ))}
        </div>
    );
};

export default Grid;
