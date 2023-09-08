export interface CaractericasVehiculosResponse {
  id: string;
  nombre: string;
  registrado: string;
  vehiculos:number;
}

export interface CreateCaractericaVehiculos {
  id?: string;
  nombre: string;
}

export const CrearInput: CreateCaractericaVehiculos = {
  id: undefined,
  nombre: '',
}
