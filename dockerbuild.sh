# Build Backend Image
cd backend
docker build -t rtf-backend .
cd ..

# Build Frontend Image
cd client
docker build -t rtf-frontend .
cd ..
