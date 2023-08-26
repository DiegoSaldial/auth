export interface Roles {
  id: string;
  nombres: string;
  bit: number;
}

export interface RolesResponse {
  id: string;
  nombres: string;
  bit: number;
  permisos: number;
  usuarios: number;
}

export interface RolesResponseCustom {
  id: string;
  nombres: string;
  check: boolean;
}

export interface ModificarRol {
  id?: string;
  nombre: string;
  permisos: string[];
}
