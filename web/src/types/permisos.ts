export interface PermisosResponse {
  metodo: string;
  descripcion: string;
  rol_bits: number;
}

export interface PermisosResponseCustom {
  metodo: string;
  descripcion: string;
  rol_bits: number;
  check: boolean;
}
