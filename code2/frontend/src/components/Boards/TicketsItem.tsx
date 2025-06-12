import Alert from '@mui/material/Alert';

import "./Boards.css";
import { ITicket } from "../../interfaces/Ticket/ITicket";
import { ConfirmationWindow } from "../ModalWindows/ConfirmationWindow";
import { RefundIcon } from '../Icons/RefundIcon';
import { useWindow } from "../../hooks/useWindows/useWindow";


interface TicketsItemProps {
	ticket: ITicket
	ticketRefund: (ticketUid: string) => Promise<void>
}

export function TicketsItem({ ticket, ticketRefund }: TicketsItemProps) {	
	const confirmDeleteWindow = useWindow();

	return (
		<>
			<div className="bg-white rounded-lg shadow-md hover:shadow-lg transition-shadow p-6">
				<div className="flex justify-between items-start mb-4">
					<div className="text-xl font-semibold text-gray-800">
						Рейс #{ticket.flightNumber}
					</div>
					{ticket.status === "PAID" && (
						<RefundIcon 
							color="gray"
							addClassName="px-2 py-2 hover:bg-gray-900/10"
							onClick={confirmDeleteWindow.handleOpenWindow}
						/>
					)}
				</div>

				<div className="space-y-4">
					<div className="flex justify-between items-center">
						<div className="text-gray-600">Откуда</div>
						<div className="text-gray-800 font-medium">{ticket.fromAirport}</div>
					</div>

					<div className="flex justify-between items-center">
						<div className="text-gray-600">Куда</div>
						<div className="text-gray-800 font-medium">{ticket.toAirport}</div>
					</div>

					<div className="flex justify-between items-center">
						<div className="text-gray-600">Дата</div>
						<div className="text-gray-800 font-medium">{ticket.date}</div>
					</div>

					<div className="flex justify-between items-center">
						<div className="text-gray-600">Цена</div>
						<div className="text-gray-800 font-medium">{ticket.price} ₽</div>
					</div>
				</div>

				<div className="mt-6">
					{ticket.status === "PAID" ? (
						<Alert
							sx={{ fontSize: 16 }}
							severity="success"
						>
							Билет оплачен
						</Alert>
					) : (
						<Alert
							sx={{ fontSize: 16 }}
							severity="warning"
						>
							Билет сдан
						</Alert>
					)}
				</div>
			</div>

			{confirmDeleteWindow.visibility && (
				<ConfirmationWindow 
					header="Подтвердите возврат билета"
					onClose={confirmDeleteWindow.handleCloseWindow}
					onConfirm={async () => {
						await ticketRefund(ticket.ticketUid);
						confirmDeleteWindow.handleCloseWindow();
					}}
				>
					<div className="space-y-4">
						<div className="flex justify-between items-center">
							<div className="text-gray-600">Номер рейса</div>
							<div className="text-gray-800 font-medium">{ticket.flightNumber}</div>
						</div>

						<div className="flex justify-between items-center">
							<div className="text-gray-600">Откуда</div>
							<div className="text-gray-800 font-medium">{ticket.fromAirport}</div>
						</div>

						<div className="flex justify-between items-center">
							<div className="text-gray-600">Куда</div>
							<div className="text-gray-800 font-medium">{ticket.toAirport}</div>
						</div>

						<div className="flex justify-between items-center">
							<div className="text-gray-600">Дата</div>
							<div className="text-gray-800 font-medium">{ticket.date}</div>
						</div>

						<div className="flex justify-between items-center">
							<div className="text-gray-600">Цена</div>
							<div className="text-gray-800 font-medium">{ticket.price} ₽</div>
						</div>
					</div>
				</ConfirmationWindow>
			)}
		</>
	)
}
