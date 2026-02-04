import { Moon, Sun } from 'lucide-react';
import { useUIStore } from '../../store/uiStore';

export function ThemeToggle() {
  const { darkMode, toggleDarkMode } = useUIStore();

  return (
    <button
      onClick={toggleDarkMode}
      className="w-full justify-start px-3 py-2 text-sm font-medium rounded-md hover:bg-accent/60 text-muted-foreground hover:text-foreground transition-colors duration-150 flex items-center gap-2.5"
      title={darkMode ? 'Switch to light mode' : 'Switch to dark mode'}
    >
      {darkMode ? (
        <Sun className="w-4 h-4 flex-shrink-0" />
      ) : (
        <Moon className="w-4 h-4 flex-shrink-0" />
      )}
      {darkMode ? 'Light Mode' : 'Dark Mode'}
    </button>
  );
}
