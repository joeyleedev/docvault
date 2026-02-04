import { useUIStore } from '../../store/uiStore';
import { SidebarHeader } from './SidebarHeader';
import { DocumentList } from '../documents/DocumentList';
import { Button } from '../ui/button';
import { ThemeToggle } from '../ui/ThemeToggle';
import { FilePlus, Menu, Plus, Moon, Sun } from 'lucide-react';

export function Sidebar() {
  const { sidebarCollapsed, toggleSidebar, setCreateDialogOpen, darkMode, toggleDarkMode } = useUIStore();

  return (
    <aside
      className={`
        fixed left-0 top-0 bottom-0 z-50
        flex glass-sidebar transition-all duration-300 ease-out
        ${sidebarCollapsed ? 'w-14' : 'w-72'}
      `}
    >
      {sidebarCollapsed ? (
        <div className="flex flex-col items-center py-3 w-full h-full">
          <button
            onClick={toggleSidebar}
            className="p-2 rounded-md hover:bg-accent/60 text-muted-foreground hover:text-foreground transition-colors duration-150"
            title="Expand sidebar"
          >
            <Menu className="w-5 h-5" />
          </button>
          <div className="flex-1" />
          <button
            onClick={toggleDarkMode}
            className="p-2 rounded-md hover:bg-accent/60 text-muted-foreground hover:text-foreground transition-colors duration-150"
            title={darkMode ? 'Switch to light mode' : 'Switch to dark mode'}
          >
            {darkMode ? <Sun className="w-5 h-5" /> : <Moon className="w-5 h-5" />}
          </button>
          <button
            onClick={() => setCreateDialogOpen(true)}
            className="p-2 rounded-md hover:bg-accent/60 text-muted-foreground hover:text-foreground transition-colors duration-150"
            title="New document"
          >
            <Plus className="w-5 h-5" />
          </button>
        </div>
      ) : (
        <div className="flex flex-col w-full h-full">
          <SidebarHeader />
          <div className="flex-1 overflow-hidden px-2">
            <DocumentList />
          </div>
          <div className="p-3 border-t border-white/20 dark:border-white/10 space-y-1">
            <Button
              onClick={() => setCreateDialogOpen(true)}
              className="w-full justify-start"
              variant="ghost"
            >
              <FilePlus className="w-4 h-4 mr-2" />
              New Document
            </Button>
            <ThemeToggle />
          </div>
        </div>
      )}
    </aside>
  );
}
