export interface CategoriaVehiculosResponse {
  id: string;
  nombre: string;
  descripcion: string;
  vehiculos: number;
}

export interface CreateCategoriaVehiculos {
  id?: string;
  nombre: string;
  descripcion: string;
}

export const CrearInput: CreateCategoriaVehiculos = {
  id: undefined,
  nombre: '',
  descripcion: '',
}
