<?php

namespace App\Controller\Api;


use App\Entity\Product;
use App\Repository\ProductRepository;
use Doctrine\ORM\EntityManagerInterface;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\Routing\Annotation\Route;

#[Route('/api/products')]
final class ProductController extends AbstractController
{
    #[Route('', name: 'api_product_index', methods: ['GET'])]
    public function index(ProductRepository $productRepository): JsonResponse
    {
        $products = $productRepository->findAll();
        $data = [];

        foreach ($products as $p) {
            $data[] = [
                'id' => $p->getId(),
                'name' => $p->getName(),
                'price' => $p->getPrice(),
                'description' => $p->getDescription(),
            ];
        }

        return $this->json($data);
    }

    #[Route('', name: 'api_product_post', methods: ['POST'])]
    public function post(Request $req, EntityManagerInterface $entityManager): JsonResponse
    {
        $data = json_decode($req->getContent(), true);
        if (!$data || !isset($data['name']) || !isset($data['price'])) {
            return $this->json(['message' => 'Request must contain (name, price)'], 400);
        }

        $p = new Product();
        $p->setName($data['name']);
        $p->setPrice((float) $data['price']);
        $p->setDescription($data['description'] ?? null);

        $entityManager->persist($p);
        $entityManager->flush();

        return $this->json(['message' => 'Product created', 'id' => $p->getId()], 201);
    }
        
    #[Route('/{id}', name: 'api_product_get', methods: ['GET'])]
    public function get(int $id, ProductRepository $productRepository): JsonResponse
    {
        $p = $productRepository->find($id);
        if (!$p) {
            return $this->json(['message' => 'Product not found'], 404);
        }

        return $this->json([
            'id' => $p->getId(),
            'name' => $p->getName(),
            'price' => $p->getPrice(),
            'description' => $p->getDescription(),
        ]);
    }

    #[Route('/{id}', name: 'api_product_edit', methods: ['PUT', 'PATCH'])]
    public function edit(int $id, Request $req, ProductRepository $productRepository, EntityManagerInterface $entityManager): JsonResponse
    {
        $p = $productRepository->find($id);

        if (!$p) {
            return $this->json(['message' => 'Product not found'], 404);
        }

        $data = json_decode($req->getContent(), true);

        if (isset($data['name'])) {
            $p->setName($data['name']);
        }
        if (isset($data['price'])) {
            $p->setPrice((float) $data['price']);
        }
        if (isset($data['description'])) {
            $p->setDescription($data['description']);
        }

        $entityManager->flush();

        return $this->json(['message' => 'Product updated']);
    }

    #[Route('/{id}', name: 'api_product_delete', methods: ['DELETE'])]
    public function delete(int $id, ProductRepository $productRepository, EntityManagerInterface $entityManager): JsonResponse
    {
        $p = $productRepository->find($id);

        if (!$p) {
            return $this->json(['message' => 'Product not found'], 404);
        }

        $entityManager->remove($p);
        $entityManager->flush();

        return $this->json(['message' => 'Product deleted']);
    }    
}
