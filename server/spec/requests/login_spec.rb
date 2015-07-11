require 'rails_helper'

RSpec.describe 'Login', type: :request do

  describe 'POST /login' do
    let(:user) { create(:user, :with_stupid_password) }

    context 'with a valid request' do
      before do
        post login_path, name: user.name, password: 'secret'
      end

      it 'should work' do
        expect(response).to have_http_status(200)
      end
      it 'should return a reasponse that includes the user' do
        expect(response_json[:user][:id]).to eq user.id
      end
      it 'should return a reasponse that includes the session cookie' do
        expect(response.headers['Set-Cookie']).to include('submarine_api_session')
      end
    end

    context 'with invalid params' do
      it 'should not work' do
        expect { post login_path }.to raise_error(Virtus::CoercionError)
      end
    end

    context 'with an incorrect password' do
      it 'should raise login error' do
        expect { post login_path, name: user.name, password: 'incorrect' }.to raise_error(ApplicationError::Login)
      end
    end

    context 'with an incorrect user name' do
      it 'should raise login error' do
        expect { post login_path, name: 'unknown', password: 'secret' }.to raise_error(ApplicationError::Login)
      end
    end

  end

end
