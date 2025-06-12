import { IFlight } from "../../../interfaces/Flight/IFlight";
import { ITicketResponse } from "../../../interfaces/Ticket/ITicketResponse";
import { IUser } from '../../../interfaces/User/IUser';
import { IPrivilege } from "../../../interfaces/Bonus/IPrivilege";
import { BuyTicketWindow } from "../../ModalWindows/BuyTicketWindow";
import { useWindow } from "../../../hooks/useWindows/useWindow";


interface FlightsCardProps {
	flight: IFlight
	user: IUser | null
	privilege: IPrivilege | null
	handleOpenPurchaseInfoWindow: (ticket: ITicketResponse) => void
	handleUpdatePrivilege: () => Promise<void>
}

export function FlightsCard(props: FlightsCardProps) {
	const buyTicketWindow = useWindow();

	return (
		<>
			<div className="bg-white rounded-lg shadow-md hover:shadow-lg transition-shadow p-6">
				<div className="flex justify-between items-start mb-4">
					<div className="text-xl font-semibold text-gray-800">
						Рейс #{props.flight.flightNumber}
					</div>
					<div className="text-2xl font-bold text-blue-600">
						{props.flight.price} ₽
					</div>
				</div>

				<div className="space-y-4">
					<div className="flex justify-between items-center">
						<div className="text-gray-600">Откуда</div>
						<div className="text-gray-800 font-medium">{props.flight.fromAirport}</div>
					</div>

					<div className="flex justify-between items-center">
						<div className="text-gray-600">Куда</div>
						<div className="text-gray-800 font-medium">{props.flight.toAirport}</div>
					</div>

					<div className="flex justify-between items-center">
						<div className="text-gray-600">Дата</div>
						<div className="text-gray-800 font-medium">{props.flight.date}</div>
					</div>
				</div>

				{props.user && (
					<div className="mt-6">
						<button
							onClick={buyTicketWindow.handleOpenWindow}
							className="w-full py-2 px-4 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
						>
							Купить билет
						</button>
					</div>
				)}
			</div>

			{buyTicketWindow.visibility && (
				<BuyTicketWindow
					flight={props.flight}
					privilege={props.privilege}
					onClose={buyTicketWindow.handleCloseWindow}
					handleOpenPurchaseInfoWindow={props.handleOpenPurchaseInfoWindow}
					handleUpdatePrivilege={props.handleUpdatePrivilege}
				/>
			)}
		</>
	);
} 