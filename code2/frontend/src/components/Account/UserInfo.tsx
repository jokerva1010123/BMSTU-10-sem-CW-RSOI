import "./Account.css";
import { IUser } from "../../interfaces/User/IUser";
import { TextRow } from "../Texts/TextRow";


interface UserInfoProps {
	user: IUser
}

export function UserInfo({ user }: UserInfoProps) {
	return (
		<div className="bg-white rounded-lg shadow-md p-6 mb-6">
			<div className="text-2xl font-semibold text-gray-700 mb-4">Информация о пользователе</div>
			<div className="grid grid-cols-1 md:grid-cols-2 gap-4">
				<TextRow 
					label="Логин"
					text={ user.login }
				/>
				<TextRow 
					label="Имя"
					text={ user.firstname }
				/>
				<TextRow 
					label="Фамилия"
					text={ user.lastname }
				/>
				<TextRow 
					label="Роль"
					text={ user.role }
				/>
				<TextRow 
					label="Почта"
					text={ user.email }
				/>
				<TextRow 
					label="Телефон"
					text={ user.phone }
				/>
			</div>
		</div>
	)
}
