import "../Tables.css";
import { TablePagination } from '../TablePagination';
import { DataLoadError } from "../../DataLoadError/DataLoadError";
import { FilterFlightsWindow } from "../../ModalWindows/FilterFlightsWindow";
import { PurchaseInfoWindow } from "../../ModalWindows/PurchaseInfoWindow";
import { ITicketResponse } from "../../../interfaces/Ticket/ITicketResponse";
import { IUser } from '../../../interfaces/User/IUser';
import { usePurchaseInfoWindow } from "../../../hooks/useWindows/usePurchaseInfoWindow";
import { useFlightsTable } from "../../../hooks/useTables/useFlightsTable";
import { useFilterFlightsWindow } from "../../../hooks/useWindows/useFilterFlightsWindow";
import { FlightsCard } from "./FlightsCard";
import { InputRow } from "../../Inputs/InputRow";
import { DateTimeSelection } from "../../Selects/DateTimeSelection";
import { FormButton } from "../../Buttons/FormButton";
import { FilterIcon } from "../../Icons/FilterIcon";
import dayjs, { Dayjs } from 'dayjs';


interface FlightsTableProps {
	openMiniDrawer: boolean
	user: IUser | null
}

export function FlightsTable({ openMiniDrawer, user }: FlightsTableProps) {
	const { 
		privilege,
		flights,
		amountFlights,
		sortTable, 
		filterTable,
		page, 
		rowsPerPage,
		error,
		handleUpdatePrivilege,
		handleUpdateTable,
		handleChangePage, 
		handleChangeRowsPerPage,
		handleChangeSort,
		handleChangeFilter,
	} = useFlightsTable();

	const filterFlightsWindow = useFilterFlightsWindow({ handleChangeFilter });
	const purchaseInfoWindow = usePurchaseInfoWindow();

	const handleFilter = (e: React.FormEvent) => {
		e.preventDefault();
		handleChangeFilter(filterTable);
	};

	const clearFilter = () => {
		handleChangeFilter({
			flightNumber: "",
			fromAirport: "",
			toAirport: "",
			minDate: null,
			maxDate: null,
			minPrice: 0,
			maxPrice: 0
		});
	};

	const getDateValue = (date: Dayjs | null | undefined): Dayjs | null => {
		return date || null;
	};

	return (
		<>
			<div className={`${openMiniDrawer ? "short-table-container" : "long-table-container"}`}>
				<div className="table">
					<div className="flex justify-between items-center mb-4">
						<div className="text-2xl font-semibold text-gray-700">Список полетов</div>
					</div>

					{/* Expanded Filter Section */}
					<div className="bg-white rounded-lg shadow-sm p-4 mb-4">
						<form onSubmit={handleFilter} className="space-y-3">
							{/* First Row */}
							<div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-3">
								<InputRow
									label="Номер полета"
									value={filterTable.flightNumber || ""}
									setValue={(value) => handleChangeFilter({...filterTable, flightNumber: value})}
								/>
								<InputRow
									label="Аэропорт отправления"
									value={filterTable.fromAirport || ""}
									setValue={(value) => handleChangeFilter({...filterTable, fromAirport: value})}
								/>
								<InputRow
									label="Аэропорт прибытия"
									value={filterTable.toAirport || ""}
									setValue={(value) => handleChangeFilter({...filterTable, toAirport: value})}
								/>
							</div>

							{/* Second Row */}
							<div className="grid grid-cols-2 md:grid-cols-4 gap-3">
								<DateTimeSelection 
									label="Время вылета"
									value={getDateValue(filterTable.minDate)}
									setValue={(value) => handleChangeFilter({...filterTable, minDate: value})}
									addClassName="w-full"
								/>
								<DateTimeSelection 
									label="Время прилета"
									value={getDateValue(filterTable.maxDate)}
									setValue={(value) => handleChangeFilter({...filterTable, maxDate: value})}
									addClassName="w-full"
								/>
								<div className="flex items-end space-x-2">
									<InputRow
										label="Цена от"
										value={(filterTable.minPrice || 0).toString()}
										setValue={(value) => handleChangeFilter({...filterTable, minPrice: Number(value)})}
									/>
									<InputRow
										label="Цена до"
										value={(filterTable.maxPrice || 0).toString()}
										setValue={(value) => handleChangeFilter({...filterTable, maxPrice: Number(value)})}
									/>
								</div>
							</div>

							<div className="flex justify-end space-x-2">
								<button
									type="button"
									onClick={clearFilter}
									className="flex items-center px-3 py-1.5 text-sm text-gray-600 hover:bg-gray-100 rounded-lg transition-colors"
								>
									<FilterIcon
										color="gray"
										selected={false}
										addClassName="mr-1"
									/>
									Очистить
								</button>
								<FormButton text="Применить" />
							</div>
						</form>
					</div>

					<div className="rows-container">
						{ !error
							?	<div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
									{ flights.map((flight) => 
											<FlightsCard 
												key={ flight.flightNumber }
												flight={ flight } 
												user={ user }
												privilege={ privilege }
												handleOpenPurchaseInfoWindow={ purchaseInfoWindow.handleOpenWindow }
												handleUpdatePrivilege={ handleUpdatePrivilege }
											/>
										)
									}
								</div>
							: <DataLoadError 
									handleUpdate={ handleUpdateTable }
								/>
						}
					</div>

					<TablePagination
						amountItems={ amountFlights }
						page={ page }
						handleChangePage={ handleChangePage }
						rowsPerPage={ rowsPerPage }
						handleChangeRowsPerPage={ handleChangeRowsPerPage }
					/>
				</div>
			</div>

			{ purchaseInfoWindow.visibility && 
				<PurchaseInfoWindow 
					ticket={ purchaseInfoWindow.ticket as ITicketResponse }
					onClose={ purchaseInfoWindow.handleCloseWindow }
				/>
			}
		</>
	)
}
