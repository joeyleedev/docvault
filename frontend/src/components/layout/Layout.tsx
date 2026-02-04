import { useEffect } from 'react';
import { Outlet } from 'react-router-dom';
import { Sidebar } from './Sidebar';
import { Toaster } from '../ui/sonner';
import { useUIStore } from '../../store/uiStore';
import { CreateDialog } from '../dialogs/CreateDialog';
import { clsx } from 'clsx';

export function Layout() {
  const { toggleSidebar, setCreateDialogOpen, createDialogOpen, sidebarCollapsed } = useUIStore();

  useEffect(() => {
    const handleKeyDown = (e: KeyboardEvent) => {
      const isMac = navigator.platform.toUpperCase().indexOf('MAC') >= 0;
      const modifierKey = isMac ? e.metaKey : e.ctrlKey;

      if (modifierKey) {
        switch (e.key.toLowerCase()) {
          case 'b':
            e.preventDefault();
            toggleSidebar();
            break;
          case 'n':
            e.preventDefault();
            setCreateDialogOpen(true);
            break;
        }
      }
    };

    window.addEventListener('keydown', handleKeyDown);

    return () => {
      window.removeEventListener('keydown', handleKeyDown);
    };
  }, [toggleSidebar, setCreateDialogOpen]);

  return (
    <div className="flex h-screen bg-background text-foreground">
      <Sidebar />
      <main
        className={clsx(
          'flex-1 overflow-hidden transition-all duration-300 ease-out',
          sidebarCollapsed ? 'ml-14' : 'ml-72'
        )}
      >
        <Outlet />
      </main>
      <CreateDialog
        isOpen={createDialogOpen}
        onClose={() => setCreateDialogOpen(false)}
      />
      <Toaster />
    </div>
  );
}
