import { useState } from 'react';
import '../styles/header.css';

const Header = () => {
    const [isMenuOpen, setIsMenuOpen] = useState(false);

    return (
        <header className="header">

            <div className="logo-section">
                <div className="logo">
                    MyClearHome
                    <br></br>
                    #МойЧистыйДом
                </div>
            </div>

            <nav className={`nav ${isMenuOpen ? 'nav-open' : ''}`}>
                <a href="#main">Главная</a>
                <a href="#services">Услуги</a>
                <a href="#prices">Цены</a>
                <a href="#contacts">Контакты</a>
                <div className="header-actions">
                    <div className="phone-number">
                        <span>+7 (999) 123-45-67</span>
                    </div>
                    <button className="cta-button">
                        Заказать уборку
                    </button>
                </div>
            </nav>


            <button
                className="menu-toggle"
                onClick={() => setIsMenuOpen(!isMenuOpen)}
                aria-label="Открыть меню"
            >

            </button>

        </header>
    );
};

export default Header;