from typing import List, Union, Tuple

import torch
from transformers import AutoModelForSequenceClassification, AutoTokenizer


class Ranker:
    def __init__(self, pretrained: str, device: str, batch_size: int = 1):
        self.tokenizer = AutoTokenizer.from_pretrained(pretrained)
        self.model = AutoModelForSequenceClassification.from_pretrained(pretrained)
        self.model.to(device)
        self.model.eval()
        self.batch_size: int = batch_size

    def rank(self, query: str, docs: List[str], return_scores: bool = False) -> List[Union[Tuple[int, float], int]]:
        scores: List[float] = []
        for i in range(0, len(docs), self.batch_size):
            batch = [(query, candidate) for candidate in docs[i: i + self.batch_size]]
            with torch.no_grad():
                inputs = self.tokenizer(batch, padding=True, truncation=True, return_tensors='pt', max_length=512)
                scores.extend(self.model(**inputs, return_dict=True).logits.view(-1, ).float().detach().cpu().tolist())
        sorted_idx_list = sorted(range(len(docs)), key=lambda x: scores[x], reverse=True)
        return [(i, scores[i]) if return_scores else i for i in sorted_idx_list]
