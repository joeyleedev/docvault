import { create } from 'zustand';
import { persist } from 'zustand/middleware';

// Get initial dark mode from localStorage or system preference
const getInitialDarkMode = (): boolean => {
  const stored = localStorage.getItem('docvault-dark-mode');
  if (stored !== null) {
    return JSON.parse(stored);
  }
  return window.matchMedia('(prefers-color-scheme: dark)').matches;
};

interface UIState {
  sidebarCollapsed: boolean;
  toggleSidebar: () => void;
  setSidebarCollapsed: (collapsed: boolean) => void;
  createDialogOpen: boolean;
  setCreateDialogOpen: (open: boolean) => void;
  darkMode: boolean;
  toggleDarkMode: () => void;
  setDarkMode: (mode: boolean) => void;
}

export const useUIStore = create<UIState>()(
  persist(
    (set) => ({
      sidebarCollapsed: false,
      toggleSidebar: () => set((state) => ({ sidebarCollapsed: !state.sidebarCollapsed })),
      setSidebarCollapsed: (collapsed) => set({ sidebarCollapsed: collapsed }),
      createDialogOpen: false,
      setCreateDialogOpen: (open) => set({ createDialogOpen: open }),
      darkMode: getInitialDarkMode(),
      toggleDarkMode: () => set((state) => ({ darkMode: !state.darkMode })),
      setDarkMode: (mode) => set({ darkMode: mode }),
    }),
    {
      name: 'docvault-ui-storage',
      partialize: (state) => ({
        sidebarCollapsed: state.sidebarCollapsed,
        darkMode: state.darkMode,
      }),
    }
  )
);
