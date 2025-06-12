import Alert from '@mui/material/Alert';
import AccountBalanceWalletIcon from '@mui/icons-material/AccountBalanceWallet';

import "./Boards.css";
import { TicketsItem } from "./TicketsItem";
import { DataLoadError } from "../DataLoadError/DataLoadError";
import { IUser } from '../../interfaces/User/IUser';
import { useTicketsBoard } from "../../hooks/useBoard/useTicketsBoard";


interface TicketsBoardProps {
	openMiniDrawer: boolean
	user: IUser
}

export function TicketsBoard({ openMiniDrawer, user }: TicketsBoardProps) {
	const { 
		userInfo,
		error,
		handleUpdateTable,
		ticketRefund,
	} = useTicketsBoard();

	return (
		<>
			<div className={`${openMiniDrawer ? "short-board-container" : "long-board-container"}`}>
				<div className="board">
					{userInfo && (
						<div className="bg-white rounded-lg shadow-md p-8 mb-6">
							<div className="flex items-center space-x-6">
								<div className="bg-blue-50 p-4 rounded-full">
									<AccountBalanceWalletIcon className="text-blue-600" style={{ fontSize: 36 }} />
								</div>
								<div>
									<div className="text-lg text-gray-500">Баланс бонусов</div>
									<div className="text-4xl font-bold text-gray-800">
										{userInfo.privilege.balance}
									</div>
								</div>
							</div>
							<div className="mt-6 text-lg text-gray-600">
								Привет, {user.firstname}! Используйте бонусы для получения скидок при покупке билетов.
							</div>
						</div>
					)}

					<div className="flex justify-between items-center mb-6">
						<div className="text-2xl font-semibold text-gray-700">Мои билеты</div>
					</div>

					{!error ? (
						userInfo?.tickets && userInfo.tickets.length > 0 ? (
							<div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
								{userInfo.tickets.map(ticket => (
									<TicketsItem 
										ticket={ticket}
										ticketRefund={ticketRefund}
										key={ticket.ticketUid} 
									/>
								))}
							</div>
						) : (
							<div className="text-center text-xl text-gray-600">Нет билетов</div>
						)
					) : (
						<DataLoadError 
							handleUpdate={handleUpdateTable}
						/>
					)}
				</div>
			</div>
		</>
	)
}
